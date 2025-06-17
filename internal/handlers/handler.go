package handlers

import (
	"context"
	"fmt"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"foxminded/3.3-weather-forecast-bot/external/holiday"
	"foxminded/3.3-weather-forecast-bot/external/holiday/holidayModels"
	"foxminded/3.3-weather-forecast-bot/external/location"
	"foxminded/3.3-weather-forecast-bot/external/weather"
	"foxminded/3.3-weather-forecast-bot/internal/models"
	"foxminded/3.3-weather-forecast-bot/internal/services/subscription"
	"foxminded/3.3-weather-forecast-bot/internal/state"
	"foxminded/3.3-weather-forecast-bot/slogger"
)

type HandlerI interface {
	HandleCommand(ctx context.Context, update tgbotapi.Update) tgbotapi.MessageConfig
	HandleCallback(ctx context.Context, update tgbotapi.Update) (tgbotapi.Chattable, error)
}

type Handler struct {
	holiday holiday.ServiceI
	weather weather.ServiceI
	state   state.StorageI
	service subscription.ServiceI
}

func NewHandler(holiday holiday.ServiceI, weather weather.ServiceI, location *location.MockServiceI, state state.StorageI, service subscription.ServiceI) *Handler {
	return &Handler{holiday: holiday, weather: weather, state: state, service: service}
}

// handleCommand processes incoming commands
func (h *Handler) HandleCommand(ctx context.Context, update tgbotapi.Update) tgbotapi.MessageConfig {

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

	switch {
	case update.Message.IsCommand():
		switch update.Message.Command() {
		case models.CommandStart, models.CommandHelp:
			msg.Text = models.Start
			return msg
		case models.CommandAbout:
			msg.Text = models.About
			return msg
		case models.CommandLinks:
			msg.Text = models.Links
			return msg
		case models.CommandHoliday:
			msg.Text = models.MsgSelectCountry
			msg.ReplyMarkup = models.GetKeyboard()
			return msg
		case models.CommandWeather:
			msg.Text = models.MsgSentLocation
			msg.ReplyMarkup = models.GetKeyboardLocationReq()
			return msg
		case models.CommandSubscribe:
			h.state.Set(update.Message.Chat.ID, state.StateWaitingLocation)
			msg.Text = models.MsgSentLocation
			msg.ReplyMarkup = models.GetKeyboardLocationReq()
			return msg
		case models.CommandSubscriptions:
			keyboard, err := h.service.GetSubsByUserAndKeyboard(ctx, update.Message.Chat.ID)
			if err != nil {
				slogger.Log.ErrorContext(ctx, "failed get subs by user and keyboard", "err", err)
				msg.Text = "Failed to get subscriptions."
				return msg
			}

			if len(keyboard.InlineKeyboard) == 0 {
				msg.Text = "No subscriptions found."
				return msg
			}

			msg.Text = "🗑 Tap on a subscription below to delete it:"
			msg.ReplyMarkup = keyboard
			return msg
		case models.CommandUnsubscribeAll:
			err := h.service.DeleteSubscriptions(ctx, update.Message.Chat.ID)
			if err != nil {
				slogger.Log.ErrorContext(ctx, "failed to delete subscriptions", "err", err)
				msg.Text = "Failed to delete subscriptions."
				return msg
			}
			msg.Text = "Subscription unsubscribed"
			return msg
		default:
			msg.Text = models.DefaultMsg
			return msg

		}

	case update.Message.Location != nil:
		userState := h.state.Get(update.Message.Chat.ID)

		if userState == state.StateWaitingLocation {
			h.state.Set(update.Message.Chat.ID, state.StateWaitingTime)

			h.service.SetUserLocation(update.Message.Chat.ID, update.Message.Location.Latitude, update.Message.Location.Longitude)

			msg.Text = "⏰ Enter time (HH:MM) for daily updates:"
			return msg
		}

		w, err := h.weather.GetWeather(ctx, update.Message.Location.Latitude, update.Message.Location.Longitude)
		msg.Text = w
		if err != nil {
			slogger.Log.ErrorContext(ctx, "Failed to get weather", "err", err)
			msg.Text = models.MsgWeatherErr
		}

		msg.ParseMode = tgbotapi.ModeMarkdownV2
		return msg

	default:
		userState := h.state.Get(update.Message.Chat.ID)

		if userState == state.StateWaitingTime {
			err := h.service.CreateSubscription(ctx, update.Message.Chat.ID, update.Message.Text)
			if err != nil {
				slogger.Log.ErrorContext(ctx, "Failed to create subscription", "err", err)
				msg.Text = "Failed to create subscription."
				return msg
			}

			h.state.Clear(update.Message.Chat.ID)
			msg.Text = "✅ You've been subscribed to daily updates."
			return msg
		}
		if countryCode, exists := holidayModels.CountryFlags[update.Message.Text]; exists {
			holidays, err := h.holiday.GetHolidayToday(ctx, countryCode)
			if err != nil {
				slogger.Log.ErrorContext(ctx, "Failed to get holidays", "err", err)
				msg.Text = models.MsgHolidayErr
				return msg
			}
			msg.Text = holidays
			return msg
		} else {
			msg.Text = models.DefaultMsg
			return msg
		}
	}
}

func (h *Handler) HandleCallback(ctx context.Context, update tgbotapi.Update) (tgbotapi.Chattable, error) {
	data := update.CallbackQuery.Data
	chatID := update.CallbackQuery.Message.Chat.ID

	if strings.HasPrefix(data, "unsubscribe:") {
		id := strings.TrimPrefix(data, "unsubscribe:")
		err := h.service.DeleteSubscription(ctx, id)
		if err != nil {
			return nil, fmt.Errorf("failed to delete subscription: %w", err)
		}
		edit := tgbotapi.NewEditMessageText(chatID, update.CallbackQuery.Message.MessageID, "Subscription removed")
		return edit, nil
	}
	return tgbotapi.NewMessage(chatID, "Unknown command."), fmt.Errorf("unknown command")
}
