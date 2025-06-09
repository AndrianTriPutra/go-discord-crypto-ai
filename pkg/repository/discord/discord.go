package discord

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

type repo struct {
	s         *discordgo.Session
	channelID map[string]string
	msg1      chan []byte
	msg2      chan []byte
	msg3      chan []byte
}

type RepositoryI interface {
	Close() error
	Send(channelID, data string) error
	SendImage(channelID, filePath string) error

	OnGocap() chan []byte
	OnTrade() chan []byte
	OnSniper() chan []byte
}

func NewRepo(token string, channelID map[string]string) (RepositoryI, error) {
	msg1 := make(chan []byte)
	msg2 := make(chan []byte)
	msg3 := make(chan []byte)

	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, fmt.Errorf("New:%w", err)
	}

	dg.Identify.Intents =
		discordgo.IntentsGuildMessages |
			discordgo.IntentMessageContent

	r := &repo{
		channelID: channelID,
		s:         dg,
		msg1:      msg1,
		msg2:      msg2,
		msg3:      msg3,
	}
	dg.AddHandler(r.handler)

	err = dg.Open()
	if err != nil {
		return nil, fmt.Errorf("Open:%w", err)
	}

	return r, nil
}

func (r *repo) Close() error {
	return r.s.Close()
}

func (r *repo) OnGocap() chan []byte {
	return r.msg1
}

func (r *repo) OnTrade() chan []byte {
	return r.msg2
}

func (r *repo) OnSniper() chan []byte {
	return r.msg3
}
