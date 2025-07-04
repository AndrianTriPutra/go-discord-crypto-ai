package dex_test

import (
	"encoding/json"
	"fmt"
	"godibot-atp/pkg/repository/dex"
	"godibot-atp/pkg/utils/logger"
	"testing"
)

func Test_Dex_1F(t *testing.T) {
	repo := dex.NewRepo()
	tokens, err := repo.Trending()
	if err != nil {
		t.Fatalf("Trending: %v", err)
	}
	if tokens == nil || len(*tokens) == 0 {
		logger.Level("fatal", "Test", "expected non-empty token list, got nil or empty")
		t.Fatal("expected non-empty token list, got nil or empty")
	}
	for i, token := range *tokens {
		js, _ := json.MarshalIndent(token, "", "  ")
		logger.Level("info", "Test", fmt.Sprintf("[%d]Token: %s", i, js))
	}
}

func Test_Dex_1J(t *testing.T) {
	repo := dex.NewRepo()
	pairs, err := repo.Search("")
	if err != nil {
		t.Fatalf("Search: %v", err)
	}
	if pairs == nil || len(*pairs) == 0 {
		t.Fatal("expected non-empty pairs, got nil or empty")
	}

	for _, token := range *pairs {
		js, _ := json.MarshalIndent(token, "", "  ")
		logger.Level("info", "Test", fmt.Sprintf("pair: %s", js))
	}
}
