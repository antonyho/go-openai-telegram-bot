package gpt_test

import (
	"testing"

	"github.com/antonyho/go-openai-telegram-bot/pkg/gpt"
	"github.com/stretchr/testify/assert"
)

func TestBot_Start_emptyToken(t *testing.T) {
	bot := gpt.NewBot("")
	err := bot.Start()

	assert.Error(t, err)

	// Output:
	// Fail to start bot. Not Found
}

func TestWebhookBot_Start_emptyToken(t *testing.T) {
	webhookBot := gpt.NewWebhookBot("", "example.com")
	err := webhookBot.Start()

	assert.Error(t, err)

	// Output:
	// Fail to start bot. Not Found
}

func TestWebhookBot_Start_emptyWebhookURL(t *testing.T) {
	webhookBot := gpt.NewWebhookBot("dummytoken", "")
	err := webhookBot.Start()

	assert.Error(t, err)

	// Output:
	// FWebhook URL cannot be empty. empty URL
}
