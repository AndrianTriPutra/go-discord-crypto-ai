package domain

import "time"

type Candle struct {
	Timestamp time.Time
	Time      int64
	Open      float64
	High      float64
	Low       float64
	Close     float64
	Volume    float64
	Turnover  float64
}

type CandlesResponse struct {
	Code string     `json:"code"`
	Data [][]string `json:"data"`
}

type TAReq struct {
	Ticker string `json:"ticker"`
	TF     string `json:"tf"`
}
