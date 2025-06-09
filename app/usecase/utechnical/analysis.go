package utechnical

import (
	"context"
	"errors"
	"fmt"
	"godibot-atp/pkg/utils/domain"
	"godibot-atp/pkg/utils/logger"
	"strings"
)

func (u *usecase) Analysis(ctx context.Context, channelID string, candles []domain.Candle) error {
	//calculate full array
	closes := make([]float64, len(candles))
	for i, c := range candles {
		closes[i] = c.Close
	}

	ema, err := u.repoTA.EMA(closes)
	if err != nil {
		return fmt.Errorf("ema:%w", err)
	}
	// logger.Trace("[Analysis]", fmt.Sprintf("EMA : \n%s", ema))
	// logger.Trace("====================", "====================")
	if err := u.repoDisc.Send(channelID, ema); err != nil {
		logger.Level("error", "Analysis", "send ema")
	}

	macd, err := u.repoTA.MACD(closes)
	if err != nil {
		return fmt.Errorf("macd:%w", err)
	}
	// logger.Trace("[Analysis]", fmt.Sprintf("MACD : \n%s", macd))
	// logger.Trace("====================", "====================")
	if err := u.repoDisc.Send(channelID, macd); err != nil {
		logger.Level("error", "Analysis", "send macd")
	}

	rsi, fastK, fastD := u.repoTA.SRSI(closes)
	n := len(fastK)
	if n < 2 {
		return errors.New("data not enough for srsi")
	}

	k := fastK[n-1]
	d := fastD[n-1]
	prevK := fastK[n-2]
	prevD := fastD[n-2]

	signal := u.repoTA.SignalWithThreshold(k, d, prevK, prevD)
	// logger.Trace("[Analysis]", fmt.Sprintf("Signal RSI&SRSI: %s", signal))
	// logger.Trace("====================", "====================")
	if err := u.repoDisc.Send(channelID, signal); err != nil {
		logger.Level("error", "Analysis", "send signal")
	}

	divergence := u.repoTA.Divergence(candles, rsi)
	// logger.Trace("[Analysis]", fmt.Sprintf("Divergence : \n%s", divergence))
	// logger.Trace("====================", "====================")
	if strings.Contains(divergence, ">") {
		if err := u.repoDisc.Send(channelID, divergence); err != nil {
			logger.Level("error", "Analysis", "send divergence")
		}
	}

	fr := u.repoTA.Fibonacci_Retracement(candles)
	//logger.Trace("[Analysis]", fmt.Sprintf("Fibonacci_Retracement : \n%s", fr))
	if err := u.repoDisc.Send(channelID, fr); err != nil {
		logger.Level("error", "Analysis", "send fr")
	}

	return nil
}
