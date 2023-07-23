package main

import (
	"fmt"
	"os"

	gochatgptsdk "github.com/ak9024/go-chatgpt-sdk"
)

func main() {
	c := gochatgptsdk.NewConfig(gochatgptsdk.Config{
		OpenAIKey: os.Getenv("OPENAI_API_KEY"),
	})

	resp, _ := c.ImagesVariations(gochatgptsdk.ModelImagesVariations{
		Image: "./assets/example-img.png",
	})

	fmt.Println(resp)
}
