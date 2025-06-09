package qai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"godibot-atp/pkg/utils/logger"
	"io"
	"net/http"
	"time"
)

type Setting struct {
	Model   string
	Host    string
	Timeout time.Duration
}

type repo struct {
	setting Setting
}

type requestPayload struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
	Stream bool   `json:"stream"`
}

type RepositoryI interface {
	Think(prompt string) (string, error)
}

func NewRepo(setting Setting) RepositoryI {
	return &repo{
		setting: setting,
	}
}

func (r *repo) Think(prompt string) (string, error) {
	payload := requestPayload{
		Model:  r.setting.Model,
		Prompt: prompt,
		Stream: false,
	}
	body, _ := json.Marshal(payload)

	client := &http.Client{Timeout: r.setting.Timeout}
	resp, err := client.Post(r.setting.Host, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return "post", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Sprintf("status code [%d][%s]", resp.StatusCode, resp.Status), err
	}

	respBody, _ := io.ReadAll(resp.Body)
	var result map[string]interface{}
	if err := json.Unmarshal(respBody, &result); err != nil {
		return "decode response", err
	}

	// logger.Trace("[Think] respBody:", string(respBody))
	logger.Trace("[Think] Response:", result["response"])
	response, ok := result["response"].(string)
	if !ok {
		return "no response !ok:", nil
	}

	return response, nil
}
