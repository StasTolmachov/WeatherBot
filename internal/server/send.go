package server

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type SendMessageI interface {
	SendMessage(token string, chatID int64, text string) error
}

type SendMessage struct {
}

func NewSendMessage() *SendMessage {
	return &SendMessage{}
}

func (s *SendMessage) SendMessage(token string, chatID int64, text string) error {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return fmt.Errorf("failed to create bot api: %w", err)
	}

	msg := tgbotapi.NewMessage(chatID, text)
	msg.ParseMode = tgbotapi.ModeMarkdownV2

	_, err = bot.Send(msg)
	if err != nil {
		return fmt.Errorf("failed to send message: %w", err)
	}
	return nil
}
