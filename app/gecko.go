package app

import (
	"context"
	"encoding/json"
	"fmt"
	"godibot-atp/pkg/utils/domain"
	"godibot-atp/pkg/utils/logger"
	"strings"
)

func (a *apps) gecko(ctx context.Context, msg []byte, payload domain.Payload) error {
	content := strings.ReplaceAll(payload.Content, "#gecko", "")
	content = strings.TrimSpace(content)
	logger.Trace("[gecko] content:", string(content))

	var req domain.GeckoRequest
	if err := json.Unmarshal([]byte(content), &req); err != nil {
		logger.Level("error", "gecko", fmt.Sprintf("[gecko] content:%s->%s", string(msg), err.Error()))

		if err := a.repoDisc.Send(payload.ChannelID, fmt.Sprintf("[gecko] Unmarshal content->%s", err.Error())); err != nil {
			return fmt.Errorf("Send:%w", err)
		}

		return nil
	}

	if err := a.uGecko.CoinGecko(ctx, payload.ChannelID, &req); err != nil {
		return fmt.Errorf("CoinGecko:%w", err)
	}

	return nil
}
