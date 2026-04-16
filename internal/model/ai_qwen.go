package model

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"os"
)

type QwenClient struct {
	apiKey string
	model  string
}

func NewQwenClient() *QwenClient {
	return &QwenClient{
		apiKey: os.Getenv("DASHSCOPE_API_KEY"),
		model:  "qwen-max",
	}
}

func (c *QwenClient) Chat(ctx context.Context, prompt string) (string, error) {
	url := "https://dashscope.aliyuncs.com/api/v1/services/aigc/text-generation/generation"

	reqBody := map[string]interface{}{
		"model": c.model,
		"input": map[string]interface{}{
			"messages": []map[string]string{
				{"role": "user", "content": prompt},
			},
		},
		"parameters": map[string]interface{}{
			"result_format": "message",
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

	output, ok := result["output"].(map[string]interface{})
	if !ok {
		return "", nil
	}

	choices, ok := output["choices"].([]interface{})
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

func (c *QwenClient) GetModelName() string {
	return c.model
}