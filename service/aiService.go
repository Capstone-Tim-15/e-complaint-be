package service

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/sashabaranov/go-openai"
)

type AIService interface {
	ResolveComplaint(ctx echo.Context, complaint string) (string, error)
}

type AIServiceImpl struct {
	Client *openai.Client
}

func NewAIService(Client *openai.Client) *AIServiceImpl {
	return &AIServiceImpl{Client: Client}
}

func (c *AIServiceImpl) ResolveComplaint(ctx echo.Context, complaint string) (string, error) {
	resp, err := c.Client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: "You are a helpful assistant that resolves customer complaints.",
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: complaint,
				},
			},
		},
	)

	if err != nil {
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}
