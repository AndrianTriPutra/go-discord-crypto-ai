package discord_test

import (
	"godibot-atp/pkg/repository/discord"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

func Test_SendText_1A(t *testing.T) {
	// ctx := context.Background()
	token := ""
	channelID :=
		map[string]string{
			"general":   "",
			"go-cap":    "",
			"go-trade":  "",
			"go-sniper": "",
		}

	repo, err := discord.NewRepo(token, channelID)
	if err != nil {
		t.Fatalf("NewRepo: %v", err)
	}

	if err := repo.Send(channelID["general"], "test"); err != nil {
		t.Fatalf("Send: %v", err)
	}
}

func Test_SendImg_1A(t *testing.T) {
	// ctx := context.Background()
	token := ""
	channelID :=
		map[string]string{
			"general":   "",
			"go-cap":    "",
			"go-trade":  "",
			"go-sniper": "",
		}

	repo, err := discord.NewRepo(token, channelID)
	if err != nil {
		t.Fatalf("NewRepo: %v", err)
	}

	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)
	base := basepath[0:strings.Index(basepath, "pkg")]
	path := base + ".chart/BTC-USDT.png"

	if err := repo.SendImage(channelID["go-trade"], path); err != nil {
		t.Fatalf("SendImage: %v", err)
	}
}
