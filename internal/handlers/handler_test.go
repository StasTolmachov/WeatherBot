package handlers

import (
	"context"
	"errors"
	"testing"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	holidayMocks "foxminded/3.3-weather-forecast-bot/external/holiday"
	locationMocks "foxminded/3.3-weather-forecast-bot/external/location"
	weatherMocks "foxminded/3.3-weather-forecast-bot/external/weather"
	"foxminded/3.3-weather-forecast-bot/internal/models"
	subsMocks "foxminded/3.3-weather-forecast-bot/internal/services/subscription"
	"foxminded/3.3-weather-forecast-bot/internal/state"
	"foxminded/3.3-weather-forecast-bot/slogger"
)

func Test_HandleCommand(t *testing.T) {

	tests := []struct {
		name         string
		update       tgbotapi.Update
		expectMsg    string
		expectMarkup interface{}
		InitState    string
	}{
		{
			name:      models.CommandStart,
			update:    newUpdateWithCommand(models.CommandStart),
			expectMsg: models.Start,
		},
		{
			name:      models.CommandHelp,
			update:    newUpdateWithCommand(models.CommandHelp),
			expectMsg: models.Start,
		},
		{
			name:      models.CommandAbout,
			update:    newUpdateWithCommand(models.CommandAbout),
			expectMsg: models.About,
		},
		{
			name:      models.CommandLinks,
			update:    newUpdateWithCommand(models.CommandLinks),
			expectMsg: models.Links,
		},
		{
			name:         models.CommandHoliday,
			update:       newUpdateWithCommand(models.CommandHoliday),
			expectMsg:    models.MsgSelectCountry,
			expectMarkup: models.GetKeyboard(),
		},
		{
			name:      "GetHolidayToday",
			update:    newUpdateWithText("🇺🇸"),
			expectMsg: "Mocked Holiday",
			InitState: state.StateNone,
		},
		{
			name:      models.DefaultMsg,
			update:    newUpdateWithCommand(""),
			expectMsg: models.DefaultMsg,
		},
		{
			name:         "GetWeather_KeyboardLocationReq",
			update:       newUpdateWithCommand(models.CommandWeather),
			expectMsg:    models.MsgSentLocation,
			expectMarkup: models.GetKeyboardLocationReq(),
		},
		{
			name:      "GetWeather",
			update:    newUpdateWithLocation(2.22, 3.33),
			expectMsg: "Mocked Weather",
			InitState: state.StateNone,
		},
		{
			name:         "Subscribe_KeyboardLocationReq",
			update:       newUpdateWithCommand(models.CommandSubscribe),
			expectMarkup: models.GetKeyboardLocationReq(),
			expectMsg:    models.MsgSentLocation,
		},
		{
			name:         "Subscribes_KeyboardListOfSubscriptionsReq",
			update:       newUpdateWithCommand(models.CommandSubscriptions),
			expectMarkup: InlineKeyboardMarkup(),
			expectMsg:    "🗑 Tap on a subscription below to delete it:",
		},
		{
			name:      models.CommandUnsubscribeAll,
			update:    newUpdateWithCommand(models.CommandUnsubscribeAll),
			expectMsg: "Subscription unsubscribed",
		},
		{
			name:      "Subscribe_StateWaitingLocation",
			update:    newUpdateWithLocation(2.22, 3.33),
			expectMsg: "⏰ Enter time (HH:MM) for daily updates:",
			InitState: state.StateWaitingLocation,
		},
		{
			name:      "Subscribe_StateWaitingTime",
			update:    newUpdateWithText(""),
			expectMsg: "✅ You've been subscribed to daily updates.",
			InitState: state.StateWaitingTime,
		},
	}

	//Run custom slogger
	slogger.MakeLogger(true)

	ctx := context.Background()

	holiday := holidayMocks.NewMockServiceI(t)
	holiday.On("GetHolidayToday", mock.Anything, "US").Return("Mocked Holiday", nil)

	weather := weatherMocks.NewMockServiceI(t)
	weather.On("GetWeather", mock.Anything, 2.22, 3.33).Return("Mocked Weather", nil)

	location := locationMocks.NewMockServiceI(t)

	stateInMemory := state.NewMemoryStorage()

	service := subsMocks.NewMockServiceI(t)
	service.On("SetUserLocation", mock.Anything, mock.Anything, mock.Anything).Return().Maybe()
	service.On("GetSubsByUserAndKeyboard", mock.Anything, int64(1)).Return(InlineKeyboardMarkup(), nil)
	service.On("DeleteSubscriptions", mock.Anything, int64(1)).Return(nil)
	service.On("CreateSubscription", mock.Anything, mock.Anything, mock.Anything).Return(nil).Maybe()

	h := NewHandler(holiday, weather, location, stateInMemory, service)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if tt.InitState != "" {
				stateInMemory.Set(1, tt.InitState)
			}
			got := h.HandleCommand(ctx, tt.update)
			assert.Equal(t, tt.expectMsg, got.Text)

			if tt.expectMarkup != nil {
				assert.Equal(t, tt.expectMarkup, got.ReplyMarkup)
			} else {
				assert.Nil(t, got.ReplyMarkup)
			}
		})
	}

}

func Test_HandleCommandError(t *testing.T) {

	tests := []struct {
		name      string
		update    tgbotapi.Update
		expectMsg string
		state     string
	}{

		{
			name:      "GetHolidayTodayWithError",
			update:    newUpdateWithText("🇺🇸"),
			expectMsg: models.MsgHolidayErr,
		},
		{
			name:      "GetWeatherWithError",
			update:    newUpdateWithLocation(2.22, 3.33),
			expectMsg: models.MsgWeatherErr,
		},
		{
			name:      "SubscriptionsCommandError",
			update:    newUpdateWithCommand(models.CommandSubscriptions),
			expectMsg: "Failed to get subscriptions.",
		},
		{
			name:      "UnsubscribeAllCommandError",
			update:    newUpdateWithCommand(models.CommandUnsubscribeAll),
			expectMsg: "Failed to delete subscriptions.",
		},
		{
			name:      "SubscribeCommand_StateWaitingTimeError",
			update:    newUpdateWithText(""),
			expectMsg: "Failed to create subscription.",
			state:     state.StateWaitingTime,
		},
	}

	//Run custom slogger
	slogger.MakeLogger(true)

	ctx := context.Background()

	holidayErr := holidayMocks.NewMockServiceI(t)
	holidayErr.On("GetHolidayToday", mock.Anything, "US").Return("", errors.New("error from mocked GetHolidayToday"))

	weatherErr := weatherMocks.NewMockServiceI(t)
	weatherErr.On("GetWeather", mock.Anything, 2.22, 3.33).Return("", errors.New("err from mocked GetWeather"))

	locationErr := locationMocks.NewMockServiceI(t)

	stateInMemory := state.NewMemoryStorage()

	service := subsMocks.NewMockServiceI(t)
	service.On("SetUserLocation", mock.Anything, mock.Anything, mock.Anything).Return().Maybe()
	service.On("GetSubsByUserAndKeyboard", mock.Anything, int64(1)).Return(InlineKeyboardMarkup(), errors.New("err from mocked GetSubsByUserAndKeyboard")).Maybe()
	service.On("DeleteSubscriptions", mock.Anything, int64(1)).Return(errors.New("err from mocked DeleteSubscriptions")).Maybe()
	service.On("CreateSubscription", mock.Anything, mock.Anything, mock.Anything).Return(errors.New("err from mocked CreateSubscription")).Maybe()

	h := NewHandler(holidayErr, weatherErr, locationErr, stateInMemory, service)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if tt.state != "" {
				stateInMemory.Set(1, tt.state)
			}
			got := h.HandleCommand(ctx, tt.update)
			assert.Equal(t, tt.expectMsg, got.Text)

		})
	}

}

func Test_HandleCallback(t *testing.T) {

	ctx := context.Background()

	service := subsMocks.NewMockServiceI(t)
	service.On("DeleteSubscription", mock.Anything, "1").Return(nil).Maybe()

	h := NewHandler(nil, nil, nil, nil, service)

	t.Run("HandleCallback_unsubscribe", func(t *testing.T) {
		expectedMsg := tgbotapi.NewEditMessageText(1, 1, "Subscription removed")
		update := newUpdateWithTextCallback("unsubscribe:1")
		result, _ := h.HandleCallback(ctx, update)
		gotMsg, ok := result.(tgbotapi.EditMessageTextConfig)
		assert.True(t, ok)
		assert.Equal(t, expectedMsg.ChatID, gotMsg.ChatID)
		assert.Equal(t, expectedMsg.MessageID, gotMsg.MessageID)
		assert.Equal(t, expectedMsg.Text, gotMsg.Text)

	})

	t.Run("HandleCallback_UnknownCommand", func(t *testing.T) {
		expectedMsg := tgbotapi.NewMessage(1, "Unknown command.")
		update := newUpdateWithTextCallback("UnknownCommand:")
		result, _ := h.HandleCallback(ctx, update)
		msg, ok := result.(tgbotapi.MessageConfig)
		assert.True(t, ok)
		assert.Equal(t, expectedMsg.ChatID, msg.ChatID)
		assert.Equal(t, expectedMsg.Text, msg.Text)

	})

}

func Test_HandleCallbackError(t *testing.T) {

	ctx := context.Background()

	service := subsMocks.NewMockServiceI(t)
	service.On("DeleteSubscription", mock.Anything, "1").Return(errors.New("error from mocked DeleteSubscription")).Once()

	h := NewHandler(nil, nil, nil, nil, service)

	t.Run("HandleCallback_unsubscribeError", func(t *testing.T) {
		update := newUpdateWithTextCallback("unsubscribe:1")
		result, err := h.HandleCallback(ctx, update)

		assert.Nil(t, result)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to delete subscription")

	})

}

func newUpdateWithCommand(text string) tgbotapi.Update {
	update := tgbotapi.Update{Message: &tgbotapi.Message{
		Chat: &tgbotapi.Chat{ID: 1},
		Text: "/" + text,
		Entities: []tgbotapi.MessageEntity{
			{
				Type:   "bot_command",
				Length: len(text) + 1,
			},
		}}}
	return update
}
func newUpdateWithText(text string) tgbotapi.Update {
	update := tgbotapi.Update{Message: &tgbotapi.Message{
		Chat: &tgbotapi.Chat{ID: 1},
		Text: text,
	}}
	return update
}
func newUpdateWithLocation(lat, lon float64) tgbotapi.Update {
	update := tgbotapi.Update{Message: &tgbotapi.Message{
		Chat: &tgbotapi.Chat{ID: 1},
		Location: &tgbotapi.Location{
			Latitude:  lat,
			Longitude: lon,
		},
	}}
	return update
}
func newUpdateWithTextCallback(text string) tgbotapi.Update {
	update := tgbotapi.Update{
		CallbackQuery: &tgbotapi.CallbackQuery{
			Data: text,
			Message: &tgbotapi.Message{
				Chat:      &tgbotapi.Chat{ID: 1},
				MessageID: 1,
			},
		},
	}
	return update
}
func InlineKeyboardMarkup() tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.InlineKeyboardMarkup{
		InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{
			{
				tgbotapi.NewInlineKeyboardButtonData("text", "data"),
			},
		},
	}
}
