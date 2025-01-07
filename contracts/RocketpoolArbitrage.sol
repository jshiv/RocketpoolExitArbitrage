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

interface IWETH {
    function transfer(address recipient, uint256 amount) external returns (bool);

    function deposit() external payable;
    function withdraw(uint wad) external;
}

interface IRETH {
    function transfer(address recipient, uint256 amount) external returns (bool);

    function burn(uint256 _rethAmount) external;
    function getTotalCollateral() external view returns (uint256);
    function getEthValue(uint256 _rethAmount) external view returns (uint256);
}

contract RocketpoolExitArbitrage is IUniswapV3SwapCallbackReceiver {
    IWETH public constant WETH = IWETH(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2);
    IRETH public constant RETH = IRETH(0xae78736Cd615f374D3085123A210448E74Fc6393);
    address immutable rocketStorage = 0x1d8f8f00cfa6758d7bE78336684788Fb0ee0Fa46;
    bytes32 constant depositPoolKey = keccak256("contract.addressrocketDepositPool");
    
    uint160 internal constant MIN_SQRT_RATIO = 4295128739;
    uint160 internal constant MAX_SQRT_RATIO = 1461446703485210103287273052203988822378723970342;

    event Arbitrage(address indexed caller, address indexed uniswapPool, uint256 amount, uint256 profit);

    constructor() {}

    receive () external payable {}

    /// @notice Withdraws rETH from RocketPool, swaps it for ETH on Uniswap, and sends the profit to the caller.
    /// @dev The caller must approve this contract to spend rETH.
    /// @param _uniswapPool The address of the Uniswap V3 pool to trade rETH for ETH
    /// @param _amount The amount of rETH to withdraw from RocketPool
    /// @param _minProfit The minimum amount of profit to keep, otherwise revert
    function arb(address _uniswapPool, uint256 _amount, uint256 _minProfit) external {
        IUniswap(_uniswapPool).swap(
            address(this),
            false, // zeroForOne
            int256(_amount),
            MAX_SQRT_RATIO - 1,
            abi.encode(_uniswapPool)
        );

        uint256 profit = address(this).balance;
        require(profit >= _minProfit, "arbitrageWithdrawUniswap: Profit too low");
        (bool success, ) = payable(msg.sender).call{value: profit}("");
        require(success, "Transfer failed.");

        emit Arbitrage(msg.sender, _uniswapPool, _amount, profit);
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
}
