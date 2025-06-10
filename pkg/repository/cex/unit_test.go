package cex_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"godibot-atp/pkg/repository/cex"
	"godibot-atp/pkg/utils/logger"
)

func Test_LastPrice(t *testing.T) {
	repo := cex.NewRepo()
	price, err := repo.LastPrice("")
	if err != nil {
		t.Fatalf("LastPrice: %v", err)
	}
	js, _ := json.MarshalIndent(price, "", " ")
	logger.Trace("price:", string(js))
}

func Test_Candle_1(t *testing.T) {
	repo := cex.NewRepo()
	raw, err := repo.GetCandle("", "1month")
	if err != nil {
		t.Fatalf("GetCandle: %v", err)
	}
	for _, row := range raw {
		js, _ := json.MarshalIndent(row, "", " ")
		fmt.Println((string(js)))
	}
}
