package subscription

import (
	"context"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"foxminded/3.3-weather-forecast-bot/external/google"
	"foxminded/3.3-weather-forecast-bot/external/utils"
	"foxminded/3.3-weather-forecast-bot/internal/db"
	"foxminded/3.3-weather-forecast-bot/internal/models"
)

type ServiceI interface {
	CreateSubscription(ctx context.Context, userID int64, localTime string) error
	SetUserLocation(chatID int64, lat, lon float64)
	GetSubsByUserAndKeyboard(ctx context.Context, userID int64) (tgbotapi.InlineKeyboardMarkup, error)
	DeleteSubscriptions(ctx context.Context, userID int64) error
	DeleteSubscription(ctx context.Context, id string) error
}

type Service struct {
	db       db.MongoDBI
	google   google.ServiceI
	tempSubs map[int64]*models.TempSubscription
}

func NewService(db db.MongoDBI, google google.ServiceI) *Service {
	return &Service{
		db:       db,
		google:   google,
		tempSubs: make(map[int64]*models.TempSubscription),
	}
}

func (s *Service) CreateSubscription(ctx context.Context, userID int64, localTimeStr string) error {

	tz, err := s.google.GetTimeZoneName(ctx, s.tempSubs[userID].Latitude, s.tempSubs[userID].Longitude)

	if err != nil {
		return fmt.Errorf("failed to get time zone name: %w", err)
	}
	utcTime, err := utils.ConvertLocalTimeToUTC(localTimeStr, tz)
	if err != nil {
		return fmt.Errorf("failed to convert local time to UTC: %w", err)
	}

	localTime, err := utils.ConvertLocalTimeToTime(localTimeStr)
	if err != nil {
		return fmt.Errorf("failed to convert time to time: %w", err)
	}
	sub := models.Subscription{UserID: userID, Latitude: s.tempSubs[userID].Latitude, Longitude: s.tempSubs[userID].Longitude, TimeZone: tz, LocalTime: localTime, NotifyAtUTC: utcTime}

	if exist, err := s.db.Exists(ctx, sub); err != nil || exist {
		if err != nil {
			return fmt.Errorf("failed to check subscription existence: %w", err)
		}
		return nil

	}
	return s.db.CreateSubscription(ctx, sub)
}

func (s *Service) SetUserLocation(chatID int64, lat, lon float64) {
	if s.tempSubs[chatID] == nil {
		s.tempSubs[chatID] = &models.TempSubscription{}
	}
	s.tempSubs[chatID].Latitude = lat
	s.tempSubs[chatID].Longitude = lon
}

func (s *Service) GetSubsByUserAndKeyboard(ctx context.Context, userID int64) (tgbotapi.InlineKeyboardMarkup, error) {
	subscriptions, err := s.db.GetUserSubscriptions(ctx, userID)
	if err != nil {
		return tgbotapi.InlineKeyboardMarkup{}, fmt.Errorf("failed to get subscriptions: %w", err)
	}
	if len(subscriptions) == 0 {
		return tgbotapi.InlineKeyboardMarkup{}, nil
	}
	var rows [][]tgbotapi.InlineKeyboardButton

	for _, sub := range subscriptions {
		label := fmt.Sprintf("%s - %s", sub.TimeZone, sub.LocalTime.Format("15:04"))

		button := tgbotapi.NewInlineKeyboardButtonData(label, "unsubscribe:"+sub.ID.Hex())
		rows = append(rows, []tgbotapi.InlineKeyboardButton{button})
	}

	return tgbotapi.NewInlineKeyboardMarkup(rows...), nil

}

func (s *Service) DeleteSubscriptions(ctx context.Context, userID int64) error {
	err := s.db.DeleteSubscriptions(ctx, userID)
	if err != nil {
		return fmt.Errorf("failed to delete subscriptions: %w", err)
	}
	return nil
}

func (s *Service) DeleteSubscription(ctx context.Context, id string) error {
	err := s.db.DeleteSubscription(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to delete subscription: %w", err)
	}
	return nil
}
