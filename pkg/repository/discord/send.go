package discord

import (
	"fmt"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func (r *repo) Send(channelID, data string) error {
	_, err := r.s.ChannelMessageSend(channelID, data)
	return err
}

func (c *repo) SendImage(channelID, filePath string) error {
	ticker := filePath[strings.LastIndex(filePath, "/")+1 : int(len(filePath)-3)]

	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open image file: %w", err)
	}
	defer file.Close()

	msg := &discordgo.MessageSend{
		Content: ticker,
		Files: []*discordgo.File{
			{
				Name:        fmt.Sprintf("%s.png", ticker),
				ContentType: "image/png",
				Reader:      file,
			},
		},
	}

	_, err = c.s.ChannelMessageSendComplex(channelID, msg)
	if err != nil {
		return fmt.Errorf("failed to send message: %w", err)
	}

	return nil
}
