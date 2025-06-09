package usniper

import (
	"context"
	"godibot-atp/pkg/repository/dex"
	"godibot-atp/pkg/repository/discord"
	"godibot-atp/pkg/utils/domain"
)

type usecase struct {
	repoDisc discord.RepositoryI
	repoDex  dex.RepositoryI
}

func NewUsecase(repoDisc discord.RepositoryI, repoDex dex.RepositoryI) UsecaseI {
	return usecase{
		repoDisc: repoDisc,
		repoDex:  repoDex,
	}
}

type UsecaseI interface {
	Scaning(ctx context.Context, channelID string, payload *domain.DexScreening) error
}
