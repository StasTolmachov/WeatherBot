package weather

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"foxminded/3.3-weather-forecast-bot/external/location"
	"foxminded/3.3-weather-forecast-bot/external/utils"
	"foxminded/3.3-weather-forecast-bot/external/weather/weatherModels"
)

const (
	KelvinToCelsius = 273.15
)

type ServiceI interface {
	GetWeather(ctx context.Context, lat, lon float64) (string, error)
}

type Service struct {
	WeatherKey   string
	WeatherUrl   string
	Client       *http.Client
	locationName location.ServiceI
}

func NewService(WeatherKey, WeatherUrl string, client *http.Client, locationService location.ServiceI) *Service {
	return &Service{
		WeatherKey:   WeatherKey,
		WeatherUrl:   WeatherUrl,
		Client:       client,
		locationName: locationService,
	}
}

func (p *Service) GetWeather(ctx context.Context, lat, lon float64) (string, error) {
	urlWeatherRequest, err := utils.MakeUrlWeather(p.WeatherUrl, p.WeatherKey, "/data/3.0/onecall", lat, lon)
	if err != nil {
		return "", fmt.Errorf("failed to make url: %w", err)
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, urlWeatherRequest, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := p.Client.Do(req)
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

	var weather weatherModels.WeatherResponse
	err = json.Unmarshal(body, &weather)
	if err != nil {
		return "", fmt.Errorf("failed to unmarshal response body: %w", err)

	}

	locationName, err := p.locationName.GetLocationName(ctx, lat, lon)
	if err != nil {
		return "", fmt.Errorf("failed to get location name: %w", err)
	}
	if locationName == "" {
		return "", fmt.Errorf("empty location name response")
	}

	return formatWeatherResponse(weather, locationName), nil

}

func formatWeatherResponse(w weatherModels.WeatherResponse, locationName string) string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("📍 *%s*\n\n", escapeMarkdownV2(locationName)))
	sb.WriteString("📅 *7\\-day forecast:*\n\n")

	for i := 0; i < 7 && i < len(w.Daily); i++ {
		day := w.Daily[i]
		date := time.Unix(day.Dt, 0).Format("Mon")
		description := day.Weather[0].Description
		tempMin := day.Temp.Min - KelvinToCelsius
		tempMax := day.Temp.Max - KelvinToCelsius

		emoji := weatherModels.WeatherEmojis[description]
		if emoji == "" {
			emoji = "🌈"
		}

		temp := fmt.Sprintf("%+.1f°C / %+.1f°C", tempMin, tempMax)
		escapedTemp := escapeMarkdownV2(temp)

		desc := strings.ToUpper(description[:1]) + description[1:]
		escapedDesc := escapeMarkdownV2(desc)

		sb.WriteString(fmt.Sprintf(
			"*%s*  %s\n`%s`\n%s\n\n",
			escapeMarkdownV2(date),
			emoji,
			escapedTemp,
			escapedDesc,
		))
	}

	return sb.String()
}
func escapeMarkdownV2(text string) string {
	replacer := strings.NewReplacer(
		`_`, `\_`,
		`*`, `\*`,
		`[`, `\[`,
		`]`, `\]`,
		`(`, `\(`,
		`)`, `\)`,
		`~`, `\~`,
		"`", "\\`",
		`>`, `\>`,
		`#`, `\#`,
		`+`, `\+`,
		`-`, `\-`,
		`=`, `\=`,
		`|`, `\|`,
		`{`, `\{`,
		`}`, `\}`,
		`.`, `\.`,
		`!`, `\!`,
	)
	return replacer.Replace(text)
}
