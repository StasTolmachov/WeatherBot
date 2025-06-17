package holiday

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"foxminded/3.3-weather-forecast-bot/external/holiday/holidayModels"
	"foxminded/3.3-weather-forecast-bot/slogger"
)

type ServiceI interface {
	GetHolidayToday(ctx context.Context, countryCode string) (string, error)
}

type Service struct {
	holidayKey string
	holidayUrl string
	httpClient *http.Client
}

func NewService(HolidayKey, HolidayUrl string, client *http.Client) *Service {
	return &Service{holidayKey: HolidayKey, holidayUrl: HolidayUrl, httpClient: client}
}

func (h *Service) GetHolidayToday(ctx context.Context, countryCode string) (string, error) {
	timeToday := getTimeToday()

	URLHolidayRequest, err := makeURL(h.holidayUrl, h.holidayKey, countryCode, timeToday)
	if err != nil {
		return "", fmt.Errorf("makeURL: %w", err)
	}

	slogger.Log.DebugContext(ctx, "request to: ", "URLHolidayRequest", URLHolidayRequest)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, URLHolidayRequest, nil)
	if err != nil {
		return "", fmt.Errorf("http.NewRequestWithContext: %w", err)
	}
	resp, err := h.httpClient.Do(req)

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {

		return "", fmt.Errorf("unexpected response status: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	slogger.Log.DebugContext(ctx, "response body", "body", string(body))
	if err != nil {
		return "", fmt.Errorf("error read resp.Body: %w", err)
	}

	var holidays []holidayModels.Holiday

	err = json.Unmarshal(body, &holidays)
	if err != nil {
		return "", fmt.Errorf("error unmarshal resp.Body: %w", err)
	}

	if len(holidays) == 0 {
		return holidayModels.NoHolidays, nil
	}
	holidaysString := formatHolidayNames(holidays)

	return holidaysString, nil
}

func formatHolidayNames(h []holidayModels.Holiday) string {
	var listOfHolidayNames []string
	for _, h := range h {
		listOfHolidayNames = append(listOfHolidayNames, h.Name)
	}
	return strings.Join(listOfHolidayNames, "\n")
}

func getTimeToday() holidayModels.TimeToday {
	now := time.Now().UTC()

	return holidayModels.TimeToday{
		Year:  now.Year(),
		Month: int(now.Month()),
		Day:   now.Day(),
	}
}

func makeURL(HolidayUrl, HolidayKey, countryCode string, TimeToday holidayModels.TimeToday) (string, error) {
	baseURL, err := url.Parse(HolidayUrl)
	if err != nil {
		return "", fmt.Errorf("invalid holidayUrl provided: %w", err)
	}
	baseURL.Path += "/v1/"

	params := url.Values{}
	params.Add("api_key", HolidayKey)
	params.Add("country", countryCode)
	params.Add("year", strconv.Itoa(TimeToday.Year))
	params.Add("month", fmt.Sprintf("%02d", TimeToday.Month))
	params.Add("day", strconv.Itoa(TimeToday.Day))

	baseURL.RawQuery = params.Encode()

	return baseURL.String(), nil
}
