package gecko_test

import (
	"encoding/json"
	"fmt"
	"godibot-atp/pkg/repository/gecko"
	"godibot-atp/pkg/utils/logger"
	"testing"
)

func Test_GeckoInfo_1(t *testing.T) {
	repo := gecko.NewRepo()
	data, err := repo.GeckoInfo("bitcoin")
	if err != nil {
		t.Fatalf("GeckoInfo: %v", err)
	}

	js, _ := json.MarshalIndent(data, "", "  ")
	logger.Level("info", "Test", fmt.Sprintf("GeckoInfo: \n%s", js))
}

func Test_GeckoList_1D(t *testing.T) {
	repo := gecko.NewRepo()
	data, err := repo.GeckoList(5, 2)
	if err != nil {
		t.Fatalf("GeckoList: %v", err)
	}

	js, _ := json.MarshalIndent(data, "", "  ")
	logger.Level("info", "Test", fmt.Sprintf("GeckoInfo: \n%s", js))
}
