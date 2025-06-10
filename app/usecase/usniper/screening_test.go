package usniper_test

import (
	"context"
	"godibot-atp/app/usecase/usniper"
	"godibot-atp/pkg/repository/dex"
	"godibot-atp/pkg/repository/discord"
	"godibot-atp/pkg/utils/domain"
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
		t.Fatalf("discord.NewRepo: %v", err)
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
		t.Fatalf("Scaning: %v", err)
	}

}
