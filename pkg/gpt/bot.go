package gpt

import (
	"errors"
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TelegramBot interface {
	Start() error
}

type Bot struct {
	telegramBotToken string
	openAITokens     map[string]string
	Debug            bool
	lastUpdateID     int
}

func NewBot(token string) *Bot {
	bot := &Bot{
		telegramBotToken: token,
		openAITokens:     make(map[string]string),
		Debug:            false,
		lastUpdateID:     0,
	}
	return bot
}

func (gptBot *Bot) Start() error {
	bot, err := connectTelegramBot(gptBot.telegramBotToken, gptBot.Debug)
	err = pullUpdates(bot)

	return err
}

func connectTelegramBot(token string, debug bool) (bot *tgbotapi.BotAPI, err error) {
	bot, err = tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Println("Fail to start bot.", err)
		return
	}
	bot.Debug = debug

	return
}

func pullUpdates(bot *tgbotapi.BotAPI) error {
	config := tgbotapi.NewUpdate(0)
	updateCh := bot.GetUpdatesChan(config)

	for update := range updateCh {
		// Check if this update contains
		// - command
		// - message
		if update.Message.IsCommand() {
		}
	}

	return nil
}

type WebhookBot struct {
	Bot
	webhookURL string
}

func NewWebhookBot(telegramBotToken string, webhookURL string) *WebhookBot {
	return &WebhookBot{
		Bot:        Bot{telegramBotToken: telegramBotToken},
		webhookURL: webhookURL,
	}
}

func (gptBot *WebhookBot) Start() error {
	if strings.TrimSpace(gptBot.webhookURL) == "" {
		err := errors.New("empty URL")
		log.Println("Webhook URL cannot be empty.", err)
		return err
	}

	bot, err := connectTelegramBot(gptBot.telegramBotToken, gptBot.Debug)

	hook, err := tgbotapi.NewWebhook(gptBot.webhookURL + bot.Token)
	if err != nil {
		log.Println("Fail to start webhook.", err)
		return err
	}

	bot.Request(hook)

	return nil
}
