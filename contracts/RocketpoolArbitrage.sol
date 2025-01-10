pragma solidity ^0.8.19;

interface IUniswap {
    function swap(
        address recipient,
        bool zeroForOne,
        int256 amountSpecified,
        uint160 sqrtPriceLimitX96,
        bytes calldata data
    ) external returns (int256 amount0, int256 amount1);
}

interface IUniswapV3SwapCallbackReceiver {
    /// @notice Called to `msg.sender` after executing a swap via IUniswapV3Pool#swap.
    /// @dev In the implementation you must pay the pool tokens owed for the swap.
    /// The caller of this method must be checked to be a UniswapV3Pool deployed by the canonical UniswapV3Factory.
    /// amount0Delta and amount1Delta can both be 0 if no tokens were swapped.
    /// @param amount0Delta The amount of token0 that was sent (negative) or must be received (positive) by the pool by
    /// the end of the swap. If positive, the callback must send that amount of token0 to the pool.
    /// @param amount1Delta The amount of token1 that was sent (negative) or must be received (positive) by the pool by
    /// the end of the swap. If positive, the callback must send that amount of token1 to the pool.
    /// @param data Any data passed through by the caller via the IUniswapV3PoolActions#swap call
    function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes calldata data) external;
}

interface IMorphoBase {
    /// @notice Executes a flash loan.
    /// @dev Flash loans have access to the whole balance of the contract (the liquidity and deposited collateral of all
    /// markets combined, plus donations).
    /// @dev Warning: Not ERC-3156 compliant but compatibility is easily reached:
    /// - `flashFee` is zero.
    /// - `maxFlashLoan` is the token's balance of this contract.
    /// - The receiver of `assets` is the caller.
    /// @param token The token to flash loan.
    /// @param assets The amount of assets to flash loan.
    /// @param data Arbitrary data to pass to the `onMorphoFlashLoan` callback.
    function flashLoan(address token, uint256 assets, bytes calldata data) external;
}

/// @title IMorphoFlashLoanCallback
/// @notice Interface that users willing to use `flashLoan`'s callback must implement.
interface IMorphoFlashLoanCallback {
    /// @notice Callback called when a flash loan occurs.
    /// @dev The callback is called only if data is not empty.
    /// @param assets The amount of assets that was flash loaned.
    /// @param data Arbitrary data passed to the `flashLoan` function.
    function onMorphoFlashLoan(uint256 assets, bytes calldata data) external;
}

interface IWETH {
    function transfer(address recipient, uint256 amount) external returns (bool);

    function deposit() external payable;
    function withdraw(uint wad) external;
    function approve(address spender, uint256 amount) external returns (bool);
}

interface IRETH {
    function transfer(address recipient, uint256 amount) external returns (bool);
    function balanceOf(address account) external view returns (uint256);

    function burn(uint256 _rethAmount) external;
    function getTotalCollateral() external view returns (uint256);
    function getEthValue(uint256 _rethAmount) external view returns (uint256);
}

contract RocketpoolExitArbitrage is IUniswapV3SwapCallbackReceiver, IMorphoFlashLoanCallback {
    IWETH public constant WETH = IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2);
    IRETH public constant RETH = IRETH(0xae78736Cd615f374D3085123A210448E74Fc6393);
    address immutable rocketStorage = 0x1d8f8f00cfa6758d7bE78336684788Fb0ee0Fa46;
    bytes32 constant depositPoolKey = keccak256("contract.addressrocketDepositPool");
    IMorphoBase immutable Morpho = IMorphoBase(0xBBBBBbbBBb9cC5e90e3b3Af64bdAF62C37EEFFCb);
    address immutable Paraswap = 0x6A000F20005980200259B80c5102003040001068;
    
    uint160 internal constant MIN_SQRT_RATIO = 4295128739;
    uint160 internal constant MAX_SQRT_RATIO = 1461446703485210103287273052203988822378723970342;

    event Arbitrage(address indexed caller, address indexed receiver, address uniswapPool, uint256 amount, uint256 profit);

    constructor() {}

    receive () external payable {}

    /// @notice Withdraws rETH from RocketPool, swaps it for ETH on Uniswap, and sends the profit to the caller.
    /// @dev The caller must approve this contract to spend rETH.
    /// @param _uniswapPool The address of the Uniswap V3 pool to trade rETH for ETH
    /// @param _amount The amount of rETH to withdraw from RocketPool
    /// @param _minProfit The minimum amount of profit to keep, otherwise revert
    function arb(address _uniswapPool, uint256 _amount, uint256 _minProfit, address _receiver) external {
        IUniswap(_uniswapPool).swap(
            address(this),
            false, // zeroForOne
            int256(_amount),
            MAX_SQRT_RATIO - 1,
            abi.encode(_uniswapPool)
        );

        uint256 profit = address(this).balance;
        require(profit >= _minProfit, "Profit too low");
        (bool success, ) = payable(_receiver).call{value: profit}("");
        require(success, "Transfer failed.");

        emit Arbitrage(msg.sender, _receiver, _uniswapPool, _amount, profit);
    }

    // see: https://github.com/Uniswap/v3-core/blob/main/contracts/interfaces/callback/IUniswapV3SwapCallback.sol
    /// @notice Called to `msg.sender` after executing a swap via IUniswapV3Pool#swap.
    /// @dev In the implementation you must pay the pool tokens owed for the swap.
    /// The caller of this method must be checked to be a UniswapV3Pool deployed by the canonical UniswapV3Factory.
    /// amount0Delta and amount1Delta can both be 0 if no tokens were swapped.
    /// @param _amountRETHDelta The amount of rETH that was sent (negative) or must be received (positive) by the pool by
    /// the end of the swap. If positive, the callback must send that amount of token0 to the pool.
    /// @param _amountWETHDelta The amount of WETH that was sent (negative) or must be received (positive) by the pool by
    /// the end of the swap. If positive, the callback must send that amount of token1 to the pool.
    /// @param _data Any data passed through by the caller via the IUniswapV3PoolActions#swap call
    function uniswapV3SwapCallback(
        int256 _amountRETHDelta, // token0
        int256 _amountWETHDelta, // token1
        bytes calldata _data
    ) external override {
        (address uniswapPool) = abi.decode(_data, (address)); 
        require(msg.sender == uniswapPool, "uniswapV3SwapCallback: Unauthorized");
        
        // 1st: burn rETH, receive ETH
        RETH.burn(uint256(-_amountRETHDelta));

        // 2nd: deposit ETH to RocketDepositPool, keep profit in ETH
        WETH.deposit{value: uint256(_amountWETHDelta)}();

        // 3rd: Repay the pool to complete the swap
        WETH.transfer(uniswapPool, uint256(_amountWETHDelta));
    }

    function arbParaswap(uint256 _amount, bytes calldata _data, uint256 _minProfit, address _receiver) external {
        Morpho.flashLoan(address(WETH), _amount, _data);

        uint256 profit = address(this).balance;
        require(profit >= _minProfit, "Profit too low");
        (bool success, ) = payable(_receiver).call{value: profit}("");
        require(success, "Transfer failed.");

        emit Arbitrage(msg.sender, _receiver, address(Morpho), _amount, profit);
    }

    function onMorphoFlashLoan(uint256 amountWethBorrowed, bytes calldata data) external override {
        WETH.approve(0x6A000F20005980200259B80c5102003040001068, amountWethBorrowed);

        (bool success, ) = Paraswap.call(data);
        require(success, "Paraswap failed");

        uint rethBalance = RETH.balanceOf(address(this));
        RETH.burn(rethBalance);

        WETH.deposit{value: amountWethBorrowed}();
        WETH.approve(msg.sender, amountWethBorrowed);
    }
}
