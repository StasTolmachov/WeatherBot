package worker

import (
	"context"
	"net/http"
	"time"

	"foxminded/3.3-weather-forecast-bot/external/location"
	"foxminded/3.3-weather-forecast-bot/external/weather"
	"foxminded/3.3-weather-forecast-bot/internal/config"
	"foxminded/3.3-weather-forecast-bot/internal/db"
	"foxminded/3.3-weather-forecast-bot/internal/server"
	"foxminded/3.3-weather-forecast-bot/slogger"
)

type ServiceI interface {
	Run(ctx context.Context, cfg config.Config)
}
type Service struct {
	cfg     config.Config
	weather weather.ServiceI
	db      db.MongoDBI
	send    server.SendMessageI
}

func NewService(cfg config.Config, db db.MongoDBI) *Service {
	client := &http.Client{}
	locationService := location.NewService(cfg.CfgWeather.Key, cfg.CfgWeather.Url, client)
	weatherService := weather.NewService(cfg.CfgWeather.Key, cfg.CfgWeather.Url, client, locationService)
	return &Service{
		cfg:     cfg,
		weather: weatherService,
		db:      db,
		send:    server.NewSendMessage(),
	}
}
func (w *Service) Run(ctx context.Context) {

	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	for {
		<-ticker.C

		now := time.Now().UTC()

		subs, err := w.db.GetDueSubscriptions(ctx, now)
		if err != nil {
			slogger.Log.Error("Failed to get due subscriptions", "err", err)
		}

		for _, sub := range subs {
			slogger.Log.Debug("found subscription", "sub", sub)

			weatherForecast, err := w.weather.GetWeather(ctx, sub.Latitude, sub.Longitude)
			if err != nil {
				slogger.Log.Error("Failed to get weather", "err", err)
			}
			err = w.send.SendMessage(w.cfg.CfgBot.Token, sub.UserID, weatherForecast)
			if err != nil {
				slogger.Log.Error("Failed to send message to telegram", "err", err)
			}
		}
	}
}
