package cex

import "godibot-atp/pkg/utils/domain"

type repo struct {
}

type RepositoryI interface {
	LastPrice(ticker string) (*domain.CexLastPrice, error)
	GetCandle(ticker, timeframe string) ([]domain.Candle, error)
}

func NewRepo() RepositoryI {
	return repo{}
}
