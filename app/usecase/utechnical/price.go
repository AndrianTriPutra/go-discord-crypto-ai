package utechnical

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

func (u *usecase) Price(ctx context.Context, channelID, data string) error {
	ticker := strings.ToUpper(data)

	price, err := u.repoCex.LastPrice(ticker)
	if err != nil {
		if err := u.repoDisc.Send(channelID, fmt.Sprintf("[Price] LastPrice->%s", err.Error())); err != nil {
			return fmt.Errorf("[coinInfo] failed Send: error LastPrice")
		}
	}
	if price == nil {
		return errors.New("price nil")
	}

	js, _ := json.MarshalIndent(price, "", " ")
	if err := u.repoDisc.Send(channelID, string(js)); err != nil {
		return fmt.Errorf("Send LastPrice:%w", err)
	}

	return nil
}
