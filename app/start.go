package app

import (
	"context"
	"fmt"
	"godibot-atp/app/usecase/ugecko"
	"godibot-atp/app/usecase/usniper"
	"godibot-atp/app/usecase/utechnical"
	"godibot-atp/app/usecase/uthink"
	"godibot-atp/pkg/repository/cex"
	"godibot-atp/pkg/repository/chart"
	"godibot-atp/pkg/repository/dex"
	"godibot-atp/pkg/repository/discord"
	"godibot-atp/pkg/repository/gecko"
	"godibot-atp/pkg/repository/qai"
	"godibot-atp/pkg/repository/technical"
	"godibot-atp/pkg/utils/logger"
	"time"
)

func (a *apps) Start(ctx context.Context) error {
	var err error

	//repository
	repoQAI := qai.NewRepo(a.setting.Qai)
	repoGecko := gecko.NewRepo()
	repoDex := dex.NewRepo()
	repoCex := cex.NewRepo()
	repoTA := technical.NewRepo()
	repoChart := chart.NewRepo()

	a.repoDisc, err = discord.NewRepo(a.setting.Token, a.setting.ChannelID)
	if err != nil {
		return fmt.Errorf("newDiscord: %w", err)
	}
	time.Sleep(1 * time.Second)

	// usecae
	a.uThink = uthink.NewUsecase(a.repoDisc, repoQAI)
	a.uGecko = ugecko.NewUsecase(a.repoDisc, repoGecko)
	a.uSniper = usniper.NewUsecase(a.repoDisc, repoDex)
	a.uTA = utechnical.NewUsecase(a.repoDisc, repoCex, repoTA, repoChart)

	// Notify start
	err = a.repoDisc.Send(a.setting.ChannelID["general"], fmt.Sprintf("%s Start", a.setting.App))
	if err != nil {
		logger.Level("error", "Start", "failed Send->"+err.Error())
	}

	return nil
}
