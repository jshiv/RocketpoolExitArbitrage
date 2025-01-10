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
        vm.label(UNISWAP_RETH_WETH_POOL, "rETH/WETH_Pool");
        vm.label(address(rocketpoolExitArb), "RocketpoolExitArb");
        vm.label(0xBBBBBbbBBb9cC5e90e3b3Af64bdAF62C37EEFFCb, "Morpho");
        vm.label(0x6A000F20005980200259B80c5102003040001068, "Paraswap");
    }

    function testArbOnForkUniswap() public {
        uint256 ethAvailableBeforeArb = MAINNET_RETH.balance;

        // 3) Impersonate a node
        address nodeAddress = vm.envAddress("NODE_ADDRESS");
        vm.label(nodeAddress, "node address");
        vm.startPrank(nodeAddress);

        // 4) Distribute minipool
        address minipoolAddress = vm.envAddress("MINIPOOL_ADDRESS");
        vm.label(minipoolAddress, "minipool address");


        RPMinipool(minipoolAddress).distributeBalance(false);

        vm.stopPrank();

        // 5) Execute the arb
        rocketpoolExitArb.arb(
            UNISWAP_RETH_WETH_POOL,
            23897692000000000000,
            58322582921635196,
            address(this)
        );

        uint256 ethAvailableAfterArb = MAINNET_RETH.balance;


        // 6) Validate results or log final balances
        //    In a real test, you'd check the user’s ETH balance, logs, etc.
        console.log("=== Test completed. Check logs or balances for profit. ===");

        // 7) check arb contract balance
        console.log("Arb contract balance: ", address(rocketpoolExitArb).balance);
        console.log("rETH balance before arb: ", ethAvailableBeforeArb);
        console.log("rETH balance after arb: ", ethAvailableAfterArb);
        require(ethAvailableBeforeArb <= ethAvailableAfterArb, "rETH balance should never decrease");
        require(ethAvailableBeforeArb + 25 > ethAvailableAfterArb, "rETH balance should increase by over 25 wei");
    }

    function testArbOnForkParaswap() public {
        uint256 ethAvailableBeforeArb = MAINNET_RETH.balance;

        // 3) Impersonate a node
        address nodeAddress = vm.envAddress("NODE_ADDRESS");
        vm.label(nodeAddress, "node address");
        vm.startPrank(nodeAddress);

        // 4) Distribute minipool
        address minipoolAddress = vm.envAddress("MINIPOOL_ADDRESS");
        vm.label(minipoolAddress, "minipool address");


        RPMinipool(minipoolAddress).distributeBalance(false);

        vm.stopPrank();

        // 5) Execute the arb
        rocketpoolExitArb.arbParaswap(
            87145304340033234, 
            hex"5e94e28d0000000000000000000000000000000000000000000000000000000000000060000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000000000000000000000000001e0000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc2000000000000000000000000ae78736cd615f374d3085123a210448e74fc639300000000000000000000000000000000000000000000000001359a314d07eed200000000000000000000000000000000000000000000000001147faa4996c8ef00000000000000000000000000000000000000000000000001359a314d07eed22d9132db1e4346b99d63a28ed7a96dcf00000000000000000000000001496353000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000060000000000000000000000000ae78736cd615f374d3085123a210448e74fc6393000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc200000000000000000000000000000000000000000000000000000000000000640000000000000000000000000000000000000000000000000000000000000000", 
            0,
            address(this)
        );

        uint256 ethAvailableAfterArb = MAINNET_RETH.balance;


        // 6) Validate results or log final balances
        //    In a real test, you'd check the user’s ETH balance, logs, etc.
        console.log("=== Test completed. Check logs or balances for profit. ===");

        // 7) check arb contract balance
        console.log("Arb contract balance: ", address(rocketpoolExitArb).balance);
        console.log("rETH balance before arb: ", ethAvailableBeforeArb);
        console.log("rETH balance after arb: ", ethAvailableAfterArb);
        require(ethAvailableBeforeArb <= ethAvailableAfterArb, "rETH balance should never decrease");
        require(ethAvailableBeforeArb + 25 > ethAvailableAfterArb, "rETH balance should increase by over 25 wei");
    }
}
