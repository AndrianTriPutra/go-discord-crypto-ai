package uthink

import (
	"context"
	"fmt"
	"godibot-atp/pkg/repository/discord"
	"godibot-atp/pkg/repository/qai"
	"godibot-atp/pkg/utils/domain"
)

type usecase struct {
	repoDisc discord.RepositoryI
	repoQAi  qai.RepositoryI
}

func NewUsecase(repoDisc discord.RepositoryI, repoQAi qai.RepositoryI) UsecaseI {
	return usecase{
		repoDisc: repoDisc,
		repoQAi:  repoQAi,
	}
}

type UsecaseI interface {
	QAI(ctx context.Context, payload *domain.Payload) error
}

func (u usecase) QAI(ctx context.Context, payload *domain.Payload) error {
	// Think AI
	resp, err := u.repoQAi.Think(payload.Content)
	if err != nil {
		return fmt.Errorf("[Think]%s:%w", resp, err)
	}

	// logger.Level("debug", "QAI", fmt.Sprintf("response:%s", resp))
	if err = u.repoDisc.Send(payload.ChannelID, resp); err != nil {
		return fmt.Errorf("Send:%w", err)
	}
	return nil
}
