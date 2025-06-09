package domain

import "time"

type CexLastResp struct {
	Code string   `json:"code"`
	Data *CexData `json:"data"`
}

type CexData struct {
	Time        int64  `json:"time"`
	Sequence    string `json:"sequence"`
	Price       string `json:"price"`
	Size        string `json:"size"`
	BestBid     string `json:"bestBid"`
	BestBidSize string `json:"bestBidSize"`
	BestAsk     string `json:"bestAsk"`
	BestAskSize string `json:"bestAskSize"`
}

type CexLastPrice struct {
	Time  time.Time `json:"time"`
	Price string    `json:"price"`
}
