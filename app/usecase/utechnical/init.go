package utechnical

import (
	"context"
	"godibot-atp/pkg/repository/cex"
	"godibot-atp/pkg/repository/chart"
	"godibot-atp/pkg/repository/discord"
	"godibot-atp/pkg/repository/technical"
	"godibot-atp/pkg/utils/domain"
)

type usecase struct {
	repoDisc  discord.RepositoryI
	repoCex   cex.RepositoryI
	repoTA    technical.RepositoryI
	repoChart chart.RepositoryI
}

func NewUsecase(
	repoDisc discord.RepositoryI,
	repoCex cex.RepositoryI,
	repoTA technical.RepositoryI,
	repoChart chart.RepositoryI) UsecaseI {
	return &usecase{
		repoDisc:  repoDisc,
		repoCex:   repoCex,
		repoTA:    repoTA,
		repoChart: repoChart,
	}
}

type UsecaseI interface {
	Price(ctx context.Context, channelID, data string) error
	Technical_Analysis(ctx context.Context, path, channelID string, req domain.TAReq) error
	Analysis(ctx context.Context, channelID string, candles []domain.Candle) error
}
