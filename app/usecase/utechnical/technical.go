package utechnical

import (
	"context"
	"fmt"
	"godibot-atp/pkg/utils/domain"
	"os"
	"strings"
	"time"
)

func (u *usecase) Technical_Analysis(ctx context.Context, path, channelID string, req domain.TAReq) error {
	ticker := strings.ToUpper(req.Ticker)
	tf := strings.ToLower(req.TF)
	ts := time.Now().UTC()

	candles, err := u.repoCex.GetCandle(ticker, tf)
	if err != nil {
		errN := fmt.Errorf("Technical_Analysis] GetCandle->%w", err)
		if err := u.repoDisc.Send(channelID, errN.Error()); err != nil {
			return fmt.Errorf("[coinInfo] failed Send: error GetCandle")
		}
		return errN
	}

	filePath := fmt.Sprintf("%s%s.png", path, ticker)
	if err := u.repoChart.GenerateChart(ts, tf, filePath, candles); err != nil {
		if err := u.repoDisc.Send(channelID, fmt.Sprintf("[Technical_Analysis] GenerateChart->%s", err.Error())); err != nil {
			return fmt.Errorf("[coinInfo] failed Send: error GenerateChart")
		}
	}

	// send img to disc
	if err := u.repoDisc.SendImage(channelID, filePath); err != nil {
		return fmt.Errorf("SendImage:%w", err)
	}

	//remove file
	if err := os.Remove(filePath); err != nil {
		return fmt.Errorf("Remove[%s]:%w", filePath, err)
	}

	if err := u.Analysis(ctx, channelID, candles); err != nil {
		if err := u.repoDisc.Send(channelID, fmt.Sprintf("[Technical_Analysis] Analysis->%s", err.Error())); err != nil {
			return fmt.Errorf("[coinInfo] failed Send: error Analysis")
		}
	}

	return nil
}
