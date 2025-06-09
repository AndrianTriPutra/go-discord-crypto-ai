package gecko

import (
	"encoding/json"
	"errors"
	"fmt"
	"godibot-atp/pkg/utils/domain"
	"net/http"
)

type repo struct {
}

type RepositoryI interface {
	GeckoInfo(id string) (*domain.GeckoResume, error)
	GeckoList(limit, page uint) (*[]domain.GeckoList, error)
}

func NewRepo() RepositoryI {
	return repo{}
}

func (r repo) GeckoInfo(id string) (*domain.GeckoResume, error) {
	url := fmt.Sprintf("https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&ids=%s", id)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("get: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code [%d][%s]", resp.StatusCode, resp.Status)
	}

	var data []domain.GeckoResume
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, fmt.Errorf("decode response: %w", err)
	}

	if len(data) == 0 {
		return nil, errors.New("no data received")
	}

	return &data[0], nil
}

func (r repo) GeckoList(limit, page uint) (*[]domain.GeckoList, error) {
	url := fmt.Sprintf("https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&order=market_cap_desc&per_page=%d&page=%d&sparkline=false", limit, page)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("get: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code [%d][%s]", resp.StatusCode, resp.Status)
	}

	var data []domain.GeckoList
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, fmt.Errorf("decode response: %w", err)
	}

	if len(data) == 0 {
		return nil, errors.New("no data received")
	}

	return &data, nil
}
