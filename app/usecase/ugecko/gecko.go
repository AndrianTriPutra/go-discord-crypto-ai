package ugecko

import (
	"context"
	"encoding/json"
	"fmt"
	"godibot-atp/pkg/repository/discord"
	"godibot-atp/pkg/repository/gecko"
	"godibot-atp/pkg/utils/domain"
	"godibot-atp/pkg/utils/logger"
)

type usecase struct {
	repoDisc  discord.RepositoryI
	repoGecko gecko.RepositoryI
}

func NewUsecase(repoDisc discord.RepositoryI, repoGecko gecko.RepositoryI) UsecaseI {
	return usecase{
		repoDisc:  repoDisc,
		repoGecko: repoGecko,
	}
}

type UsecaseI interface {
	CoinGecko(ctx context.Context, channelID string, req *domain.GeckoRequest) error
}

func (u usecase) CoinGecko(ctx context.Context, channelID string, req *domain.GeckoRequest) error {
	logger.Trace("[CoinGecko] channelID:", "content:"+req.Content)

	switch req.Content {
	case "info":
		return u.coinInfo(channelID, req)
	case "list":
		return u.coinList(channelID, req)
	default:
		if err := u.repoDisc.Send(channelID, "content not found"); err != nil {
			return fmt.Errorf("failed Send: content not found")
		}
	}

	return nil
}

func (u usecase) coinInfo(channelID string, payload *domain.GeckoRequest) error {
	if payload.ID == nil {
		if err := u.repoDisc.Send(channelID, "[coinInfo] Please provide a coin id"); err != nil {
			return fmt.Errorf("[coinInfo] failed Send: Please provide a coin id")
		}
		return nil
	}

	data, err := u.repoGecko.GeckoInfo(*payload.ID)
	if err != nil {
		if err := u.repoDisc.Send(channelID, fmt.Sprintf("[coinInfo] GeckoInfo->%s", err.Error())); err != nil {
			return fmt.Errorf("[coinInfo] failed Send: error GeckoInfo")
		}
		return nil
	}

	js, _ := json.MarshalIndent(data, "", "  ")
	if err := u.repoDisc.Send(channelID, string(js)); err != nil {
		return fmt.Errorf("[coinInfo] failed Send: GeckoInfo")
	}

	return nil
}

func (u usecase) coinList(channelID string, payload *domain.GeckoRequest) error {
	if payload.List == nil {
		if err := u.repoDisc.Send(channelID, "[coinList] Please provide a list"); err != nil {
			return fmt.Errorf("[coinList] failed Send: Please provide a list")
		}
		return nil
	}

	if payload.List.Limit <= 0 {
		if err := u.repoDisc.Send(channelID, "[coinList] Please provide a valid limit"); err != nil {
			return fmt.Errorf("[coinList] failed Send: Please provide a valid limit")
		}
		return nil
	}

	if payload.List.Page <= 0 {
		if err := u.repoDisc.Send(channelID, "[coinList] Please provide a valid page"); err != nil {
			return fmt.Errorf("[coinList] failed Send: Please provide a valid page")
		}
		return nil
	}
	data, err := u.repoGecko.GeckoList(uint(payload.List.Limit), uint(payload.List.Page))
	if err != nil {
		if err := u.repoDisc.Send(channelID, fmt.Sprintf("[coinList] CG2->%s", err.Error())); err != nil {
			return fmt.Errorf("[coinList] failed Send: error GeckoList")
		}
		return nil
	}

	js, _ := json.MarshalIndent(data, "", "  ")
	if err := u.repoDisc.Send(channelID, string(js)); err != nil {
		return fmt.Errorf("[coinList] failed Send: GeckoList")
	}

	return nil
}
