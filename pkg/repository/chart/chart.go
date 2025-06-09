package chart

import (
	"fmt"
	"godibot-atp/pkg/utils/domain"
	"time"

	"gonum.org/v1/plot"
)

type repo struct {
}

type RepositoryI interface {
	GenerateChart(ts time.Time, tf, path string, candles []domain.Candle) error
}

func NewRepo() RepositoryI {
	return repo{}
}

func (r repo) makeTimeTicks(times []float64, tf string, count int) plot.Ticker {
	if len(times) == 0 {
		return plot.DefaultTicks{}
	}

	if len(times) <= count {
		count = len(times) - 1
	}
	step := len(times) / count

	ticks := make([]plot.Tick, 0, count+1)
	for i := 0; i <= count; i++ {
		idx := i * step
		if idx >= len(times) {
			idx = len(times) - 1
		}
		t := times[idx]
		ts := time.Unix(int64(t), 0).UTC()

		// adjust format label with tf
		var label string
		switch tf {
		case "1mon", "1month":
			label = ts.Format("Jan 2006")
		case "1week", "1day":
			label = ts.Format("Jan 02")
		case "4h", "1h":
			label = ts.Format("02 15:04")
		default:
			label = ts.Format("01-02 15:04")
		}

		ticks = append(ticks, plot.Tick{Value: t, Label: label})
	}
	return plot.ConstantTicks(ticks)
}

func (r repo) makePriceTicks(min, max float64, divisions int) plot.Ticker {
	if divisions <= 0 || max <= min {
		return plot.DefaultTicks{}
	}
	step := (max - min) / float64(divisions)

	ticks := []plot.Tick{}
	for i := 0; i <= divisions; i++ {
		val := min + step*float64(i)
		label := fmt.Sprintf("%.2f", val)
		ticks = append(ticks, plot.Tick{Value: val, Label: label})
	}
	return plot.ConstantTicks(ticks)
}
