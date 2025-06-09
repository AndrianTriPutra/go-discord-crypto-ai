package app

import (
	"context"
	"godibot-atp/pkg/utils/logger"
	"time"
)

func (a *apps) sniper(ctx context.Context) error {
	ticker := time.NewTicker(a.setting.Sniper.Schedulle)
	for {
		select {
		case <-ticker.C:
			logger.Trace("sniper", "on")
			if err := a.uSniper.Scaning(ctx, a.setting.ChannelID["go-sniper"], &a.setting.Sniper.Filter); err != nil {
				logger.Level("error", "sniper", "failed Scaning->"+err.Error())
			}

		case <-ctx.Done():
			return ctx.Err()
		}
	}
}
