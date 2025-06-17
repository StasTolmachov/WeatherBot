package google

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"foxminded/3.3-weather-forecast-bot/external/utils"
	"foxminded/3.3-weather-forecast-bot/slogger"
)

type ServiceI interface {
	GetTimeZoneName(ctx context.Context, lat float64, lon float64) (string, error)
}

type Service struct {
	Key    string
	Url    string
	Client *http.Client
}

func NewService(locationKey string, locationUrl string, client *http.Client) ServiceI {
	return &Service{
		Key:    locationKey,
		Url:    locationUrl,
		Client: client,
	}
}

type timeZone struct {
	TimeZone string `json:"timeZoneId"`
}

func (s *Service) GetTimeZoneName(ctx context.Context, lat float64, lon float64) (string, error) {
	//urlRequest, err := utils.MakeUrlWeather(s.Url, s.Key, "/geo/1.0/reverse", lat, lon)

	urlRequest, err := utils.MakeUrlGoogle(s.Url, "maps/api/timezone/json", lat, lon, s.Key, time.Now().Unix())

	if err != nil {
		return "", fmt.Errorf("failed to make url: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, urlRequest, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := s.Client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to do request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to do request: %d %s", resp.StatusCode, resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	var tz timeZone
	err = json.Unmarshal(body, &tz)
	if err != nil {
		return "", fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	slogger.Log.DebugContext(ctx, "timeZone name: ", "timeZone", tz)

	return tz.TimeZone, nil

}
