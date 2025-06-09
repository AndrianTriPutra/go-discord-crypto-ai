package app

import (
	"context"
	"encoding/json"
	"fmt"
	"godibot-atp/pkg/utils/domain"
	"godibot-atp/pkg/utils/logger"
	"strings"
)

func (a *apps) technical(ctx context.Context, msg []byte, payload domain.Payload) error {
	content := strings.ReplaceAll(payload.Content, "#ta", "")
	content = strings.TrimSpace(content)
	logger.Trace("[technical] content:", string(content))

	var req domain.TAReq
	if err := json.Unmarshal([]byte(content), &req); err != nil {
		logger.Level("error", "technical", fmt.Sprintf("[technical] content:%s->%s", string(msg), err.Error()))

		if err := a.repoDisc.Send(payload.ChannelID, fmt.Sprintf("[technical] Unmarshal content->%s", err.Error())); err != nil {
			return fmt.Errorf("Send:%w", err)
		}

		return nil
	}

	if err := a.uTA.Technical_Analysis(ctx, a.setting.Chart, payload.ChannelID, req); err != nil {
		return fmt.Errorf("Technical_Analysis:%w", err)
	}

	return nil
}
