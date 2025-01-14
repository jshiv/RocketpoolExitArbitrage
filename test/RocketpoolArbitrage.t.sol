// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.19;

import "forge-std/Test.sol";
import "forge-std/console.sol";

// Your contract
import "../contracts/RocketpoolArbitrage.sol";

// Minimal ERC20 interface
interface IERC20 {
    function balanceOf(address) external view returns (uint256);
    function approve(address spender, uint256 amount) external returns (bool);
}

interface RPMinipool {
    function distributeBalance(bool _rewardsOnly) external;
}

contract RocketpoolExitArbitrageForkTest is Test {
    // --------------------------
    //  Mainnet addresses
    // --------------------------
    address constant MAINNET_WETH  = 0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2; 
    address constant MAINNET_RETH  = 0xae78736Cd615f374D3085123A210448E74Fc6393;

    // The official Uniswap V3 rETH-WETH pool (fee 0.3% or 1%, check on chain).
    // For example, here is the 0.3% rETH-WETH pool:
    address constant UNISWAP_RETH_WETH_POOL = 0x553e9C493678d8606d6a5ba284643dB2110Df823;


    // Your arbitrage contract
    RocketpoolExitArbitrage public rocketpoolExitArb;

    receive () external payable {}

    function setUp() public {
        // 1) Create a mainnet fork at a specific block (with an RPC key you set in .env)
        vm.createSelectFork(vm.envString("MAINNET_RPC_URL"));

        // 2) Deploy your RocketpoolExitArbitrage contract
        rocketpoolExitArb = new RocketpoolExitArbitrage();

        // [Optional] Labeling for better console logs
        vm.label(MAINNET_WETH, "WETH");
        vm.label(MAINNET_RETH, "rETH");
        vm.label(address(rocketpoolExitArb), "RocketpoolExitArb");
        vm.label(0xBBBBBbbBBb9cC5e90e3b3Af64bdAF62C37EEFFCb, "Morpho");
        vm.label(0x6A000F20005980200259B80c5102003040001068, "Paraswap");
    }

    function testArbOnForkUniswap() public {
        uint256 profitBefore = address(this).balance;
        uint256 arbBefore = address(rocketpoolExitArb).balance;
        uint256 rethBefore = MAINNET_RETH.balance;

        // Impersonate node
        address nodeAddress = vm.envAddress("NODE_ADDRESS");
        vm.label(nodeAddress, "node address");
        vm.startPrank(nodeAddress);

        // 1) Distribute minipool
        address minipoolAddress = 0x0e4aAff61d11a2c36cc29a744885ab73A682EaCF;
        vm.label(minipoolAddress, "minipool address");

        RPMinipool(minipoolAddress).distributeBalance(false);

        vm.stopPrank();

        address _uniswapPool = 0xa4e0faA58465A2D369aa21B3e42d43374c6F9613;
        uint160 _sqrtPriceLimitX96 = 84072759829691353131019927552;
        uint256 _amount = 12346337220937533;
        uint256 _minProfit = 61816612301843;

        vm.label(_uniswapPool, "UniswapPool");

        // 2) Execute the arb
        rocketpoolExitArb.arb(
            _uniswapPool,
            _sqrtPriceLimitX96,
            _amount,
            _minProfit,
            address(this)
        );


        console.log("=== Test completed. Check logs or balances for profit. ===");

        console.log("Profit balance before arb: ", profitBefore);
        console.log("Profit balance after arb: ", address(this).balance);
        require(address(this).balance > profitBefore, "Profit balance should increase");
        require(address(this).balance - profitBefore >= _minProfit, "Profit balance should increase by over 25 wei");
        console.log("Arb contract balance before arb: ", arbBefore);
        console.log("Arb contract balance after arb:: ", address(rocketpoolExitArb).balance);
        console.log("rETH balance before arb: ", rethBefore);
        console.log("rETH balance after arb: ", MAINNET_RETH.balance);
        require(rethBefore <= MAINNET_RETH.balance, "rETH balance should never decrease");
        require(rethBefore + 25 > MAINNET_RETH.balance, "rETH balance should increase by over 25 wei");
    }

    function testArbOnForkParaswap() public {
        uint256 profitBefore = address(this).balance;
        uint256 arbBefore = address(rocketpoolExitArb).balance;
        uint256 rethBefore = MAINNET_RETH.balance;

        // Impersonate node
        address nodeAddress = vm.envAddress("NODE_ADDRESS");
        vm.label(nodeAddress, "node address");
        vm.startPrank(nodeAddress);

        // 1) Distribute minipool
        address minipoolAddress = vm.envAddress("MINIPOOL_ADDRESS");
        vm.label(minipoolAddress, "minipool address");


        RPMinipool(minipoolAddress).distributeBalance(false);

        vm.stopPrank();

        bytes memory _swapData = bytes("");
        uint256 _amount = 12346337220937533;
        uint256 _minProfit = 61816612301843;

        // 2) Execute the arb
        rocketpoolExitArb.arbParaswap(
            _amount, 
            _swapData, 
            _minProfit,
            address(this)
        );


        console.log("=== Test completed. Check logs or balances for profit. ===");

        console.log("Profit balance before arb: ", profitBefore);
        console.log("Profit balance after arb: ", address(this).balance);
        require(address(this).balance > profitBefore, "Profit balance should increase");
        require(address(this).balance - profitBefore > 25, "Profit balance should increase by over 25 wei");
        console.log("Arb contract balance before arb: ", arbBefore);
        console.log("Arb contract balance after arb:: ", address(rocketpoolExitArb).balance);
        console.log("rETH balance before arb: ", rethBefore);
        console.log("rETH balance after arb: ", MAINNET_RETH.balance);
        require(rethBefore <= MAINNET_RETH.balance, "rETH balance should never decrease");
        require(rethBefore + 25 > MAINNET_RETH.balance, "rETH balance should increase by over 25 wei");
    }
}
