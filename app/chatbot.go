package app

import (
	"context"
	"encoding/json"
	"fmt"
	"godibot-atp/pkg/utils/domain"
	"godibot-atp/pkg/utils/logger"
	"strings"
)

func (a *apps) chatbot(ctx context.Context) error {
	for {
		select {
		case msg := <-a.repoDisc.OnGocap():
			// logger.Trace("[chatbot] OnGocap:", string(msg))

			payload, err := a.decoder("OnGocap", msg)
			if err != nil {
				logger.Level("error", "chatbot", "[OnGocap] failed decoder->"+err.Error())
				break
			}

			// response default
			if (strings.HasPrefix(payload.Content, "hi") || strings.HasPrefix(payload.Content, "hai")) && len(payload.Content) <= 3 {
				logger.Trace("[chatbot] default  :", "match hi")
				if err := a.repoDisc.Send(payload.ChannelID, "hallo saya QAi ... 👋"); err != nil {
					logger.Level("error", "chatbot", "[OnGocap] failed Send match->"+err.Error())
				}
				break
			}

			// response AI
			if err := a.uThink.QAI(ctx, &payload); err != nil {
				if errN := a.repoDisc.Send(payload.ChannelID, err.Error()); errN != nil {
					logger.Level("error", "chatbot", "[OnGocap] failed Send Think->"+errN.Error())
				}
			}

		case msg := <-a.repoDisc.OnTrade():
			// logger.Trace("[chatbot] OnTrade:", string(msg))

			payload, err := a.decoder("OnTrade", msg)
			if err != nil {
				logger.Level("error", "chatbot", "[OnTrade] failed decoder->"+err.Error())
				break
			}

			if strings.HasPrefix(payload.Content, "#gecko") {
				if err := a.gecko(ctx, msg, payload); err != nil {
					logger.Level("error", "chatbot", "[gecko] failed->"+err.Error())
				}
				break
			} else if strings.HasPrefix(payload.Content, "#price") {
				content := strings.ReplaceAll(payload.Content, "#price", "")
				content = strings.TrimSpace(content)
				logger.Trace("[OnTrade] [price] content:", string(content))

				if err := a.uTA.Price(ctx, payload.ChannelID, content); err != nil {
					if errN := a.repoDisc.Send(payload.ChannelID, err.Error()); errN != nil {
						logger.Level("error", "chatbot", "[OnTrade] failed Send Price ->"+errN.Error())
					}
				}
				break
			} else if strings.HasPrefix(payload.Content, "#ta") {
				if err := a.technical(ctx, msg, payload); err != nil {
					logger.Level("error", "chatbot", "[technical] failed->"+err.Error())
				}
				break
			}

			if err := a.repoDisc.Send(payload.ChannelID, "Tag Not Found"); err != nil {
				logger.Level("error", "chatbot", "[OnTrade] failed Send->"+err.Error())
			}

		case msg := <-a.repoDisc.OnSniper():
			// logger.Trace("[chatbot] OnSniper:", string(msg))

			payload, err := a.decoder("OnSniper", msg)
			if err != nil {
				logger.Level("error", "chatbot", "[OnSniper] failed decoder->"+err.Error())
				break
			}

			if !strings.HasPrefix(payload.Content, "#sniper") {
				if err := a.repoDisc.Send(payload.ChannelID, "Tag not Found"); err != nil {
					logger.Level("error", "chatbot", "[OnSniper] failed Send->"+err.Error())
				}
				break
			}

			content := strings.ReplaceAll(payload.Content, "#sniper", "")
			content = strings.TrimSpace(content)
			logger.Trace("[chatbot] content:", string(content))

			var req domain.DexScreening
			if err := json.Unmarshal([]byte(content), &req); err != nil {
				logger.Level("error", "chatbot", fmt.Sprintf("[OnSniper] content:%s->%s", string(msg), err.Error()))
				if err := a.repoDisc.Send(payload.ChannelID, fmt.Sprintf("[OnSniper] Unmarshal content->%s", err.Error())); err != nil {
					logger.Level("error", "chatbot", "[OnSniper] failed Send->"+err.Error())
				}
				break
			}

			if err := a.uSniper.Scaning(ctx, payload.ChannelID, &req); err != nil {
				logger.Level("error", "chatbot", "[OnSniper] failed Scaning->"+err.Error())
			}

		case <-ctx.Done():
			a.repoDisc.Close()
			return ctx.Err()
		}
	}
}

func (a *apps) decoder(ucase string, msg []byte) (domain.Payload, error) {
	var payload domain.Payload
	if err := json.Unmarshal(msg, &payload); err != nil {
		logger.Level("error", "chatbot", fmt.Sprintf("[%s] payload:%s->%s", ucase, string(msg), err.Error()))
		if err := a.repoDisc.Send(payload.ChannelID, fmt.Sprintf("[%s] Unmarshal payload->%s", ucase, err.Error())); err != nil {
			return payload, fmt.Errorf("[%s] Send:%w", ucase, err)
		}
	}

	return payload, nil
}
