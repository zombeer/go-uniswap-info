package uniswap

import (
	"fmt"
	"time"
)

type Candle struct {
	Timestamp time.Time
	Open      float64
	High      float64
	Low       float64
	Close     float64
}

func Max(x, y float64) float64 {
	if x > y {
		return x
	}
	return y
}

func Min(x, y float64) float64 {
	if x < y {
		return x
	}
	return y
}

func TruncateTimestamp(ts uint32) time.Time {
	d := 5 * time.Minute
	truncatedTs := time.UnixMilli(int64(ts) * 1000).Truncate(d)
	return truncatedTs
}

func PricesToCandles(prices []Price) [][]string {
	if len(prices) == 0 {
		return nil
	}
	result := [][]string{}
	currentCandle := Candle{}

	for i, price := range prices {
		truncatedTs := TruncateTimestamp(price.Timestamp)
		pValue := price.Value

		if i == 0 {
			currentCandle = Candle{
				Open:      pValue,
				High:      pValue,
				Low:       pValue,
				Close:     pValue,
				Timestamp: truncatedTs}
		}

		if truncatedTs == currentCandle.Timestamp {
			currentCandle.High = Max(currentCandle.High, pValue)
			currentCandle.Low = Min(currentCandle.High, pValue)
			currentCandle.Close = pValue
			continue
		}
		if truncatedTs != currentCandle.Timestamp || i == len(prices)-1 {
			result = append(
				result,
				[]string{
					fmt.Sprintf("%d", currentCandle.Timestamp.UnixMilli()),
					fmt.Sprintf("%f", currentCandle.Open),
					fmt.Sprintf("%f", currentCandle.High),
					fmt.Sprintf("%f", currentCandle.Low),
					fmt.Sprintf("%f", currentCandle.Close),
				},
			)
			currentCandle = Candle{
				Open:      pValue,
				High:      pValue,
				Low:       pValue,
				Close:     pValue,
				Timestamp: truncatedTs}
		}
	}
	return result
}

// 1661126400000
// 1661958000000000
