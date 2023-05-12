package main

import (
	"context"
	"fmt"
	"log"

	openai "github.com/sashabaranov/go-openai"
)

func main() {

	//openapikey := os.Getenv("OPENAIAPIKEY") // "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"
	openapikey := "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" // Niekoniecznie ok
	client := openai.NewClient(openapikey)
	ctx := context.Background()

	request := openai.CompletionRequest{
		Model:       openai.GPT3TextDavinci003,
		Temperature: 0.7,
		MaxTokens:   1000,
		TopP:        1.0,
		Prompt:      "Jaka jest różnica pomiędzy prądem a napięciem",
	}

	response, err := client.CreateCompletion(ctx, request)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(response.Choices[0].Text)

}
