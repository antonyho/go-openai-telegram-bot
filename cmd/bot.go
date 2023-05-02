package cmd

import (
	"github.com/antonyho/go-openai-telegram-bot/pkg/gpt"
	"github.com/spf13/cobra"
)

const (
	BotTokenFlag   = "bot-token"
	WebhookURLFlag = "url"
	BotDebugFlag   = "debug"
)

var (
	// Used for flags.
	telegramBotToken     string
	webhookURL           string
	telegramBotDebugMode bool

	rootCmd = &cobra.Command{
		Use:   "tg-gpt-bot",
		Short: "A Telegram OpenAI GPT Bot",
		Long:  ``,
	}

	activeBotCmd = &cobra.Command{
		Use:   "persistent",
		Short: "A persistent bot",
		Long:  `A persistent bot which stays alive and actively fetches new message`,
		Run: func(cmd *cobra.Command, args []string) {
			gptBot := gpt.NewBot(telegramBotToken)
			gptBot.Debug = telegramBotDebugMode
			gptBot.Start()
		},
	}

	webhookBotCmd = &cobra.Command{
		Use:   "webhook",
		Short: "A webhook bot",
		Long:  `A webhook bot which listens to incoming request from Telegram`,
		Run: func(cmd *cobra.Command, args []string) {
			gptBot := gpt.NewWebhookBot(telegramBotToken, webhookURL)
			gptBot.Debug = telegramBotDebugMode
			gptBot.Start()
		},
	}
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&telegramBotToken, BotTokenFlag, "b", "", "token for Telegram bot")
	rootCmd.MarkPersistentFlagRequired(BotTokenFlag)

	rootCmd.PersistentFlags().BoolVarP(&telegramBotDebugMode, BotDebugFlag, "d", false, "debug mode")

	webhookBotCmd.LocalFlags().StringVarP(&webhookURL, WebhookURLFlag, "h", "", "webhook URL")
	webhookBotCmd.MarkFlagRequired(WebhookURLFlag)

	rootCmd.AddCommand(activeBotCmd)
	rootCmd.AddCommand(webhookBotCmd)
}

func Execute() error {
	return rootCmd.Execute()
}
