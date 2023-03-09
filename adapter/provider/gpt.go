package provider

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"zapgpt/internal/model"
)

func GenerateGPTText(query string) (string, error) {
	req := model.GptRequest{
		Model: "gpt-3.5-turbo",
		Messages: []model.Message{
			{
				Role:    "user",
				Content: query,
			},
		},
		MaxTokens: 150,
	}

	reqJson, err := json.Marshal(req)
	if err != nil {
		return "", err
	}

	request, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(reqJson))
	if err != nil {
		return "", err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer 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")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	respBody, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	var resp model.GptResponse
	err = json.Unmarshal(respBody, &resp)
	if err != nil {
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}
