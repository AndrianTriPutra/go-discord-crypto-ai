package technical

import "godibot-atp/pkg/utils/domain"

type repo struct {
}

type RepositoryI interface {
	SRSI(closes []float64) ([]float64, []float64, []float64)
	EMA(closes []float64) (string, error)
	Divergence(c []domain.Candle, rsi []float64) string
	SignalWithThreshold(k, d float64, prevK, prevD float64) string
	Fibonacci_Retracement(candles []domain.Candle) string
	MACD(closes []float64) (string, error)
}

func NewRepo() RepositoryI {
	return repo{}
}
