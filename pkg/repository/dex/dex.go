package dex

import "godibot-atp/pkg/utils/domain"

type Filter struct {
	ChainID     string  `json:"chainId"`
	PairAddress string  `json:"pairAddr"`
	DexID       string  `json:"dexId"`
	Listing     float32 `json:"listing"` // listing token
	Liquidity   float64 `json:"liquidity"`
	Volume      float64 `json:"volume"`
	UP          float32 `json:"up"`  // up percentage priceChange (%)
	BSR         float32 `json:"bsr"` //   buy sell ratio
}

type repo struct {
}

type RepositoryI interface {
	Tending() (*[]domain.TokenProfile, error)
	Search(addr string) (*[]domain.TokenPair, error)
}

func NewRepo() RepositoryI {
	return repo{}
}
