package location

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"foxminded/3.3-weather-forecast-bot/external/utils"
	"foxminded/3.3-weather-forecast-bot/external/weather/weatherModels"
	"foxminded/3.3-weather-forecast-bot/slogger"
)

type ServiceI interface {
	GetLocationName(ctx context.Context, lat float64, lon float64) (string, error)
}

type Service struct {
	LocationKey string
	LocationUrl string
	Client      *http.Client
}

func NewService(locationKey string, locationUrl string, client *http.Client) ServiceI {
	return &Service{
		LocationKey: locationKey,
		LocationUrl: locationUrl,
		Client:      client,
	}
}

func (l *Service) GetLocationName(ctx context.Context, lat float64, lon float64) (string, error) {
	urlLocationNameRequest, err := utils.MakeUrlWeather(l.LocationUrl, l.LocationKey, "/geo/1.0/reverse", lat, lon)
	if err != nil {
		return "", fmt.Errorf("failed to make url: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, urlLocationNameRequest, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := l.Client.Do(req)
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

	var LocationName weatherModels.GeoLocationResponse
	err = json.Unmarshal(body, &LocationName)
	if err != nil {
		return "", fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	slogger.Log.DebugContext(ctx, "location name: ", "location", LocationName)

	if len(LocationName) == 0 {
		return "", fmt.Errorf("empty location name response")
	}
	return LocationName[0].Name, nil

}
