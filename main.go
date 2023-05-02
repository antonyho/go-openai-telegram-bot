package main

import (
	"fmt"
	"os"

	"github.com/antonyho/go-openai-telegram-bot/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
