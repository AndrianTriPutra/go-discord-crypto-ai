package main

import (
	"godibot-atp/app"
	"godibot-atp/pkg/repository/qai"
	"godibot-atp/pkg/utils/domain"
	"path/filepath"
	"runtime"
	"time"
)

func setting(token string) *app.Setting {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)
	path := basepath + "/.chart/"

	return &app.Setting{
		App:   product,
		Token: token,
		ChannelID: map[string]string{
			"general":   "",
			"go-cap":    "",
			"go-trade":  "",
			"go-sniper": "",
		},

		Qai: qai.Setting{
			Model:   "gemma3:4b",
			Host:    "http://localhost:11434/api/generate",
			Timeout: 60 * time.Second,
		},
		Sniper: app.Sniper{
			Schedulle: 1 * time.Hour, // periodic for check new token
			Filter: domain.DexScreening{
				TH:        1.5, // treshold new token created
				Liquidity: "25k",
				Volume:    "15k",
				UP:        25,  //%
				BSR:       1.2, //buy sell ratio
			},
		},
		Chart: path,
	}
}
