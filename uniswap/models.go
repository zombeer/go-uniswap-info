package uniswap

type Price struct {
	BlockNumber uint64  `json:"bn"`
	Value       float64 `json:"value"`
}

type TokenInfo struct {
	Address string `json:"address"`
	Symbol  string `json:"symbol"`
}

type PairInfo struct {
	Address string    `json:"pair_address"`
	Symbols string    `json:"symbols"`
	Token0  TokenInfo `json:"token0"`
	Token1  TokenInfo `json:"token1"`
	Prices  []Price   `json:"prices"`
}
