package model

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"os"
)

type ClaudeClient struct {
	apiKey string
	model  string
}

func NewClaudeClient() *ClaudeClient {
	return &ClaudeClient{
		apiKey: os.Getenv("ANTHROPIC_API_KEY"),
		model:  "claude-3-opus-20240229",
	}
}

func (c *ClaudeClient) Chat(ctx context.Context, prompt string) (string, error) {
	url := "https://api.anthropic.com/v1/messages"

	reqBody := map[string]interface{}{
		"model": c.model,
		"messages": []map[string]string{
			{"role": "user", "content": prompt},
		},
		"max_tokens": 4096,
	}

	reqBytes, err := json.Marshal(reqBody)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(reqBytes))
	if err != nil {
		return "", err
	}

	req.Header.Set("x-api-key", c.apiKey)
	req.Header.Set("anthropic-version", "2023-06-01")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	content, ok := result["content"].([]interface{})
	if !ok || len(content) == 0 {
		return "", nil
	}

	firstContent, ok := content[0].(map[string]interface{})
	if !ok {
		return "", nil
	}

	text, ok := firstContent["text"].(string)
	if !ok {
		return "", nil
	}

	return text, nil
}

func (c *ClaudeClient) GetModelName() string {
	return c.model
}