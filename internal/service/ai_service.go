package service

import (
	"context"
	"errors"
)

var ErrAIProviderUnavailable = errors.New("AI provider unavailable")

type AIModelClient interface {
	Chat(ctx context.Context, prompt string) (string, error)
	GetModelName() string
}

type AIService interface {
	GenerateAnswers(ctx context.Context, question string, models []string) (string, map[string]ModelResponse, error)
	CallModel(ctx context.Context, model string, prompt string) ModelResponse
}

type aiService struct {
	modelClients map[string]AIModelClient
}

type ModelResponse struct {
	Content string `json:"content"`
	Model   string `json:"model"`
	Success bool   `json:"success"`
	Error   string `json:"error"`
}

func NewAIService(modelClients map[string]AIModelClient) AIService {
	return &aiService{
		modelClients: modelClients,
	}
}

func (s *aiService) GenerateAnswers(ctx context.Context, question string, models []string) (string, map[string]ModelResponse, error) {
	if len(models) == 0 {
		models = []string{"qwen", "gpt", "claude"}
	}

	prompt := s.buildPrompt(question)

	responses := make(map[string]ModelResponse)

	for _, model := range models {
		response := s.CallModel(ctx, model, prompt)
		responses[model] = response
	}

	summary := s.generateSummary(responses)

	return summary, responses, nil
}

func (s *aiService) CallModel(ctx context.Context, model string, prompt string) ModelResponse {
	client, ok := s.modelClients[model]
	if !ok {
		return ModelResponse{
			Content: "",
			Model:   model,
			Success: false,
			Error:   "model not found",
		}
	}

	content, err := client.Chat(ctx, prompt)
	if err != nil {
		return ModelResponse{
			Content: "",
			Model:   client.GetModelName(),
			Success: false,
			Error:   err.Error(),
		}
	}

	return ModelResponse{
		Content: content,
		Model:   client.GetModelName(),
		Success: true,
		Error:   "",
	}
}

func (s *aiService) buildPrompt(question string) string {
	return `你是一个专业的面试题解答助手。请针对以下面试题给出详细、准确的答案。

面试题：` + question + `

请从以下几个角度进行分析：
1. 给出完整、详细的答案
2. 提供相关的知识点解释
3. 给出面试扩展问题
4. 适当的代码示例（如果适用）

请用中文回答，答案要专业、清晰、有条理。`
}

func (s *aiService) generateSummary(responses map[string]ModelResponse) string {
	successCount := 0
	totalCount := 0

	for _, resp := range responses {
		totalCount++
		if resp.Success {
			successCount++
		}
	}

	if successCount == 0 {
		return "抱歉，所有AI模型暂时不可用，请稍后重试。"
	}

	modelNames := map[string]string{
		"qwen":   "通义千问",
		"gpt":    "ChatGPT",
		"claude": "Claude",
	}

	if successCount == totalCount {
		available := []string{}
		for model := range responses {
			if name, ok := modelNames[model]; ok {
				available = append(available, name)
			}
		}
		return "三个AI模型（" + joinStrings(available) + "）都给出了回答，您可以选择其中一个或查看综合分析。"
	}

	availableModels := []string{}
	for model, resp := range responses {
		if resp.Success {
			if name, ok := modelNames[model]; ok {
				availableModels = append(availableModels, name)
			}
		}
	}

	summary := ""
	if len(availableModels) > 0 {
		summary = "有 " + intToString(successCount) + " 个AI模型（" + joinStrings(availableModels) + "）可用。"
	}

	return summary
}

func joinStrings(strs []string) string {
	result := ""
	for i, s := range strs {
		if i > 0 {
			result += "、"
		}
		result += s
	}
	return result
}

func intToString(n int) string {
	if n == 0 {
		return "0"
	}
	if n == 1 {
		return "1"
	}
	if n == 2 {
		return "2"
	}
	if n == 3 {
		return "3"
	}
	return string(rune('0' + n))
}