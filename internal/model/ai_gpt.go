package model

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"os"
)

type GPTClient struct {
	apiKey string
	model  string
}

func NewGPTClient() *GPTClient {
	return &GPTClient{
		apiKey: os.Getenv("OPENAI_API_KEY"),
		model:  "gpt-4-turbo-preview",
	}
}

func (c *GPTClient) Chat(ctx context.Context, prompt string) (string, error) {
	url := "https://api.openai.com/v1/chat/completions"

	reqBody := map[string]interface{}{
		"model": c.model,
		"messages": []map[string]string{
			{"role": "user", "content": prompt},
		},
	}

	reqBytes, err := json.Marshal(reqBody)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(reqBytes))
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", "Bearer "+c.apiKey)
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

	choices, ok := result["choices"].([]interface{})
	if !ok || len(choices) == 0 {
		return "", nil
	}

	choice, ok := choices[0].(map[string]interface{})
	if !ok {
		return "", nil
	}

	message, ok := choice["message"].(map[string]interface{})
	if !ok {
		return "", nil
	}

	content, ok := message["content"].(string)
	if !ok {
		return "", nil
	}

	return content, nil
}

func (c *GPTClient) GetModelName() string {
	return c.model
}