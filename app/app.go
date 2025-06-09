package app

import (
	"context"
	"godibot-atp/app/usecase/ugecko"
	"godibot-atp/app/usecase/usniper"
	"godibot-atp/app/usecase/utechnical"
	"godibot-atp/app/usecase/uthink"
	"godibot-atp/pkg/repository/discord"
	"godibot-atp/pkg/repository/qai"
	"godibot-atp/pkg/utils/domain"
	"time"
)

type Sniper struct {
	Schedulle time.Duration
	Filter    domain.DexScreening
}

type Setting struct {
	App       string
	Token     string
	ChannelID map[string]string
	Qai       qai.Setting
	Sniper    Sniper
	Chart     string
}

type apps struct {
	setting  Setting
	repoDisc discord.RepositoryI
	uThink   uthink.UsecaseI
	uGecko   ugecko.UsecaseI
	uSniper  usniper.UsecaseI
	uTA      utechnical.UsecaseI
}

type Application interface {
	Start(ctx context.Context) error
	Run(ctx context.Context) error
}

func NewApp(setting Setting) Application {
	return &apps{
		setting: setting,
	}
}
