package checkers

import (
	"log"
	"os"
)

func CheckOpenAiKey(key string) {
	if len(key) <= 0 {
		log.Fatal("Please set your Open AI API Key first: \n\nbabelfish set-env openai YOUR_KEY_HERE")
		os.Exit(1)
	}
}
