package requests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Choice struct {
	FinishReason string  `json:"finish_reason"`
	Index        int32   `json:"index"`
	Message      Message `json:"message"`
}

type Usage struct {
	CompletionTokens int32 `json:"completion_tokens"`
	PromptTokens     int32 `json:"prompt_tokens"`
	TotalTokens      int32 `json:"total_tokens"`
}

type GptRequestBody struct {
	Model       string    `json:"model"`
	Messages    []Message `json:"messages"`
	Temperature float64   `json:"temperature"`
}

type GptResponse struct {
	Choices []Choice `json:"choices"`
	Created int64    `json:"created"`
	Id      string   `json:"id"`
	Model   string   `json:"model"`
	Object  string   `json:"object"`
	Usage   Usage    `json:"usage"`
}

const completionUrl string = "https://api.openai.com/v1/chat/completions"

func RequestGpt4Translation(prompt string, openAiKey string) (*GptResponse, error) {
	requestBody := GptRequestBody{
		Model: "gpt-4",
		Messages: []Message{
			{
				Role:    "user",
				Content: prompt,
			},
		},
		Temperature: 0.7,
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", completionUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", openAiKey))
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result GptResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
