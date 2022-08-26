## Analyzing some key uniswap v2 DEX stats
Just a sample WEB3 showcase.

Long story short: getting key stats from factory and pairs.

### Factory
Interface documentation [here](https://docs.uniswap.org/protocol/V2/reference/smart-contracts/factory). Key values:
| Name | Type | Description | 
|--|--|--|
| pairCount | Int | amount of pairs created by the Uniswap factory |
totalVolumeUSD | BigDecimal | all time USD volume across all pairs (USD is derived)
totalVolumeETH | BigDecimal| all time volume in ETH across all pairs (ETH is derived)
totalLiquidityUSD|BigDecimal|total liquidity across all pairs stored as a derived USD amount
totalLiquidityETH|BigDecimal|total liquidity across all pairs stored as a derived ETH amount
txCount|BigInt|all time amount of transactions across all pairs

### Pair
Interface documentation [here](https://docs.uniswap.org/protocol/V2/reference/smart-contracts/pair)
