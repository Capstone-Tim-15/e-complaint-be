package config

import (
	"fmt"
	"os"

	openai "github.com/sashabaranov/go-openai"
)

func ConnectOpenAI() *openai.Client {
	LoadEnv()

	TOKEN := os.Getenv("OPEN_AI_TOKEN")
	client := openai.NewClient(TOKEN)

	fmt.Println("Connected to Open AI")
	return client
}
