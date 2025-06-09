package domain

import "time"

type TokenProfile struct {
	ChainID     string `json:"chainId"`
	Address     string `json:"tokenAddress"`
	Description string `json:"description"`
}

type TokenSearch struct {
	Schema string      `json:"schemaVersion"`
	Pairs  []TokenPair `json:"pairs"`
}

type TokenPair struct {
	ChainID       string      `json:"chainId"`
	DexID         string      `json:"dexId"`
	PairAddress   string      `json:"pairAddress"`
	BaseToken     BaseToken   `json:"baseToken"`
	FDV           float64     `json:"fdv"`
	MarketCap     float64     `json:"marketCap"`
	PairCreatedAt int64       `json:"pairCreatedAt"`
	Listing       time.Time   `json:"listing"`
	Liquidity     Liquidity   `json:"liquidity"`
	PriceChange   PriceChange `json:"priceChange"`
	Volume        Volume      `json:"volume"`
	Txns          Txns        `json:"txns"`
}

type BaseToken struct {
	Address string `json:"address"`
	Name    string `json:"name"`
	Symbol  string `json:"symbol"`
}

type Liquidity struct {
	USD   float64 `json:"usd"`
	Base  float64 `json:"base"`
	Quote float64 `json:"quote"`
}

type PriceChange struct {
	M5  float64 `json:"m5"`
	H1  float64 `json:"h1"`
	H6  float64 `json:"h6"`
	H24 float64 `json:"h24"`
}

type Volume struct {
	H24 float64 `json:"h24"`
	H6  float64 `json:"h6"`
	H1  float64 `json:"h1"`
	M5  float64 `json:"m5"`
}

type Txns struct {
	M5 TxnSummary `json:"m5"`
	H1 TxnSummary `json:"h1"`
}

type TxnSummary struct {
	Buys  int `json:"buys"`
	Sells int `json:"sells"`
}

type DexScreening struct {
	Listing   float32 `json:"listing"` // listing tokens was x hours
	Liquidity string  `json:"liquidity"`
	Volume    string  `json:"volume"`
	UP        float32 `json:"up"`  // up percentage priceChange (%)
	BSR       float32 `json:"bsr"` //   buy sell ratio
}

type DexInfo struct {
	ChainID     string    `json:"chainId"`
	DexID       string    `json:"dexId"`
	PairAddress string    `json:"pairAddress"`
	Listing     time.Time `json:"listing"`
	Liquidity   float64   `json:"liquidity-h1"`
	Volume      float64   `json:"volume-h1"`
}

type TokenWithDex struct {
	BaseToken   BaseToken `json:"baseToken"`
	Description string    `json:"description"`
	Dex         []DexInfo `json:"dex"`
}
