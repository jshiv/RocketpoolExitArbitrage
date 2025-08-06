# Threshold-Based Arbitrage Feature

This feature allows you to set a minimum profit threshold for arbitrage execution. The process will continuously monitor the expected profit and only execute the arbitrage when the profit meets or exceeds the specified threshold.

## Usage

### Basic Usage

```bash
./distribute --threshold 0.01 --minipool <minipool-address>
```

This will:
- Check the expected profit every minute (default interval)
- Wait until the profit reaches at least 0.01 ETH
- Execute the arbitrage when the threshold is met

### Advanced Usage

```bash
./distribute --threshold 0.005 --check-interval 30s --minipool <minipool-address> --protocol best
```

This will:
- Check profit every 30 seconds
- Wait for profit to reach 0.005 ETH
- Use the best protocol (Uniswap or Paraswap)
- Execute when threshold is met

## Flags

### --threshold
- **Type**: string
- **Description**: Minimum profit threshold in ETH (e.g., 0.01)
- **Example**: `--threshold 0.01`

### --check-interval
- **Type**: duration
- **Default**: 1m (1 minute)
- **Description**: Interval between profit checks
- **Examples**: 
  - `--check-interval 30s` (30 seconds)
  - `--check-interval 2m` (2 minutes)
  - `--check-interval 1h` (1 hour)

## How It Works

1. **Initialization**: The process sets up the connection and validates inputs
2. **Profit Monitoring**: Every `check-interval`, the process:
   - Calculates current expected profit from arbitrage
   - Compares it against the threshold
   - Logs the current profit vs threshold
3. **Execution**: When profit meets or exceeds the threshold:
   - Logs that threshold is met
   - Executes the normal arbitrage process
   - Sends the bundle to Flashbots

## Example Output

```
Waiting for profit to reach threshold of 0.010000 ETH...
Checking every 1m0s...

Current expected profit: 0.005000 ETH (threshold: 0.010000 ETH)
Current expected profit: 0.007500 ETH (threshold: 0.010000 ETH)
Current expected profit: 0.012000 ETH (threshold: 0.010000 ETH)
Profit threshold met! Expected profit: 0.012000 ETH

[normal arbitrage execution continues...]
```

## Use Cases

1. **Profit Optimization**: Wait for better market conditions
2. **Risk Management**: Only execute when profit justifies gas costs
3. **Automated Trading**: Set and forget arbitrage with profit targets
4. **Market Timing**: Avoid executing during unfavorable conditions

## Notes

- The process will run indefinitely until the threshold is met or the process is interrupted
- Use Ctrl+C to stop the process if needed
- All existing flags (--protocol, --receiver, etc.) work with threshold mode
- The profit calculation includes gas fees and uses the same logic as normal execution
- Threshold values are in ETH (not wei) for user convenience