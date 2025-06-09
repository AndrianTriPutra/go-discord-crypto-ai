package usniper_test

import (
	"context"
	"godibot-atp/app/usecase/usniper"
	"godibot-atp/pkg/repository/dex"
	"godibot-atp/pkg/repository/discord"
	"godibot-atp/pkg/utils/domain"
	"godibot-atp/pkg/utils/logger"
	"testing"
)

func Test_Screening_1L(t *testing.T) {
	ctx := context.Background()
	token := ""
	channelID := map[string]string{}
	channel := ""

	repoDex := dex.NewRepo()
	repoDisc, err := discord.NewRepo(token, channelID)
	if err != nil {
		logger.Level("fatal", "discord.NewRepo", err.Error())
	}

	ucase := usniper.NewUsecase(repoDisc, repoDex)

	payload := &domain.DexScreening{
		Listing: 1.5,
		// Liquidity: "50k",
		// Volume:    "15k",

		// Liquidity: "20k",
		// Volume:    "15k",

		Liquidity: "2k",
		Volume:    "1k",
		UP:        25,  // (%)
		BSR:       1.2, //   buy sell ratio
	}

	err = ucase.Scaning(ctx, channel, payload)
	if err != nil {
		logger.Level("fatal", "Scaning", err.Error())
	}

}
