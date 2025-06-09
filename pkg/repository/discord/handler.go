package discord

import (
	"encoding/json"
	"fmt"
	"godibot-atp/pkg/utils/domain"
	"godibot-atp/pkg/utils/logger"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func (r *repo) handler(s *discordgo.Session, m *discordgo.MessageCreate) {
	logger.Trace("[handler] ChannelID:", m.ChannelID)
	logger.Trace("[handler] Author   :", m.Author.ID)
	if m.Author.ID == s.State.User.ID {
		return
	}

	logger.Trace("[handler] Content  :", m.Content)
	for alias, id := range r.channelID {
		if m.ChannelID == id {
			logger.Trace("[handler] channel match:", alias)

			switch alias {
			case "go-cap":
				if strings.HasPrefix(m.Content, "#QAi") {
					content := strings.ReplaceAll(m.Content, "#QAi ", "")
					payload := &domain.Payload{
						ChannelID: m.ChannelID,
						Content:   content,
					}
					js, err := json.Marshal(payload)
					if err != nil {
						logger.Level("error", "handler", fmt.Sprintf("[go-cap] error: %s", err.Error()))
						return
					}
					r.msg1 <- js
				}

			case "go-trade":
				payload := &domain.Payload{
					ChannelID: m.ChannelID,
					Content:   m.Content,
				}
				js, err := json.Marshal(payload)
				if err != nil {
					logger.Level("error", "handler", fmt.Sprintf("[go-cap] error: %s", err.Error()))
					return
				}
				r.msg2 <- js

			case "go-sniper":
				payload := &domain.Payload{
					ChannelID: m.ChannelID,
					Content:   m.Content,
				}
				js, err := json.Marshal(payload)
				if err != nil {
					logger.Level("error", "handler", fmt.Sprintf("[go-sinper] error: %s", err.Error()))
					return
				}
				r.msg3 <- js

			default:
				logger.Level("warning", "discord", fmt.Sprintf("channel [%s] not supported", alias))
			}

			return
		}
	}
}
