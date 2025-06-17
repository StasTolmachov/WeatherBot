package server

import (
	"context"
	"fmt"
	"net/http"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/google/uuid"

	"foxminded/3.3-weather-forecast-bot/external/google"
	"foxminded/3.3-weather-forecast-bot/external/holiday"
	"foxminded/3.3-weather-forecast-bot/external/location"
	"foxminded/3.3-weather-forecast-bot/external/weather"
	"foxminded/3.3-weather-forecast-bot/internal/config"
	"foxminded/3.3-weather-forecast-bot/internal/db"
	"foxminded/3.3-weather-forecast-bot/internal/handlers"
	"foxminded/3.3-weather-forecast-bot/internal/services/subscription"
	"foxminded/3.3-weather-forecast-bot/internal/state"
	"foxminded/3.3-weather-forecast-bot/slogger"
)

type Bot struct {
	cfg     config.Config
	handler handlers.HandlerI
}

func NewBot(cfg config.Config, db db.MongoDBI) *Bot {
	client := &http.Client{}
	holidayService := holiday.NewService(cfg.CfgHoliday.Key, cfg.CfgHoliday.Url, client)
	locationService := location.NewService(cfg.CfgWeather.Key, cfg.CfgWeather.Url, client)
	googleService := google.NewService(cfg.CfgGoogle.Key, cfg.CfgGoogle.Url, client)
	weatherService := weather.NewService(cfg.CfgWeather.Key, cfg.CfgWeather.Url, client, locationService)
	memoryStorage := state.NewMemoryStorage()
	service := subscription.NewService(db, googleService)
	handler := handlers.NewHandler(holidayService, weatherService, nil, memoryStorage, service)
	return &Bot{cfg: cfg, handler: handler}
}

// RunBot starts the message processing cycle
func (b *Bot) RunBot(ctx context.Context) error {

	bot, err := tgbotapi.NewBotAPI(b.cfg.CfgBot.Token)
	if err != nil {
		return fmt.Errorf("failed to create bot api: %w", err)
	}

	slogger.Log.Info("Authorized on account:", "account", bot.Self.UserName)

	bot.Debug = b.cfg.CfgBot.Debug

	u := tgbotapi.NewUpdate(0)
	u.Timeout = b.cfg.CfgBot.Timeout

	for update := range bot.GetUpdatesChan(u) {
		if update.Message != nil {
			ctx = context.WithValue(ctx, "trace-id", uuid.New())

			msg := b.handler.HandleCommand(ctx, update)

			slogger.Log.InfoContext(ctx, "Received message", "chat_id", update.Message.Chat.ID, "user_id", update.Message.From.UserName, "text", update.Message.Text)
			slogger.Log.InfoContext(ctx, "Sending message", "chat_id", update.Message.Chat.ID, "user_id", update.Message.From.UserName, "text", msg.Text)
			if _, err := bot.Send(msg); err != nil {
				slogger.Log.ErrorContext(ctx, "error sending message", "chat_id", update.Message.Chat.ID, "user_id", update.Message.From.UserName, "text", update.Message.Text, "err", err)
			}
		} else if update.CallbackQuery != nil {
			slogger.Log.Debug("CallbackQuery")
			msg, err := b.handler.HandleCallback(ctx, update)
			if err != nil {
				slogger.Log.ErrorContext(ctx, "error HandleCallback", "err", err)
			}
			if _, err := bot.Send(msg); err != nil {
				slogger.Log.ErrorContext(ctx, "error sending message", "err", err)
			}
		}
	}
	return nil
}
