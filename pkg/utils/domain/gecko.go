package domain

import "time"

type GeckoResume struct {
	ID                    string    `json:"id"`
	Name                  string    `json:"name"`
	MarketCapRank         int       `json:"market_cap_rank"`
	MarketCap             float64   `json:"market_cap"`
	CurrentPrice          float64   `json:"current_price"`
	High24h               float64   `json:"high_24h"`
	Low24h                float64   `json:"low_24h"`
	ATH                   float64   `json:"ath"`
	ATHDate               time.Time `json:"ath_date"`
	LastUpdated           time.Time `json:"last_updated"`
	CirculatingSupply     float64   `json:"circulating_supply"`
	TotalSupply           float64   `json:"total_supply"`
	MaxSupply             *float64  `json:"max_supply"`
	FullyDilutedValuation float64   `json:"fully_diluted_valuation"`
	TotalVolume           float64   `json:"total_volume"`
}

type GeckoList struct {
	ID                string    `json:"id"`
	Name              string    `json:"name"`
	MarketCapRank     int       `json:"market_cap_rank"`
	CurrentPrice      float64   `json:"current_price"`
	CirculatingSupply float64   `json:"circulating_supply"`
	TotalSupply       float64   `json:"total_supply"`
	MaxSupply         *float64  `json:"max_supply"`
	LastUpdated       time.Time `json:"last_updated"`
}

type GeckoRequest struct {
	Content string `json:"content"`
	ID      *string `json:"id"`
	List    *struct {
		Limit int `json:"limit"`
		Page  int `json:"page"`
	} `json:"list"`
}
