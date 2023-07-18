package main

import (
	"fmt"

	gochatgptsdk "github.com/ak9024/go-chatgpt-sdk"
)

func main() {
	c := gochatgptsdk.NewConfig(gochatgptsdk.Config{
		OpenAIKey: "",
	})

	resp, _ := c.ChatCompletions(gochatgptsdk.ModelChat{
		Model: "gpt-3.5-turbo-16k",
		Messages: []gochatgptsdk.Message{
			{
				Role:    "system",
				Content: "You are a devops engineer",
			},
			{
				Role:    "user",
				Content: "Please help me to create Dockerfile for golang with a multistage image",
			},
		},
	})

	// resp, _ := c.Completions(gochatgptsdk.ModelText{
	// 	Model:  "text-davinci-003",
	// 	Prompt: "What the weather today?",
	// })

	fmt.Println(resp.Choices[len(resp.Choices)-1].Message)

	//m, _ := json.Marshal(resp)
	//fmt.Println(string(m))
}
