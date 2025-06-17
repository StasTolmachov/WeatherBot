package utils

import (
	"fmt"
	"net/url"
	"strconv"
	"time"
)

func MakeUrlWeather(apiUrl, apiKey, UrlPath string, lat, lon float64) (string, error) {
	baseUrl, err := url.Parse(apiUrl)
	if err != nil {
		return "", fmt.Errorf("invalid apiUrl provided: %w", err)
	}

	baseUrl.Path += UrlPath
	params := url.Values{}
	params.Add("lat", strconv.FormatFloat(lat, 'f', -1, 64))
	params.Add("lon", strconv.FormatFloat(lon, 'f', -1, 64))
	params.Add("appid", apiKey)

	baseUrl.RawQuery = params.Encode()

	return baseUrl.String(), nil
}

func MakeUrlGoogle(apiUrl, UrlPath string, lat, lon float64, apiKey string, timestamp int64) (string, error) {
	baseUrl, err := url.Parse(apiUrl)
	if err != nil {
		return "", fmt.Errorf("invalid apiUrl provided: %w", err)
	}

	baseUrl.Path += UrlPath
	params := url.Values{}
	location := fmt.Sprintf("%f,%f", lat, lon)
	params.Add("location", location)
	params.Add("timestamp", strconv.FormatInt(timestamp, 10))
	params.Add("key", apiKey)

	baseUrl.RawQuery = params.Encode()

	return baseUrl.String(), nil
}

func ConvertLocalTimeToUTC(input string, tz string) (time.Time, error) {
	location, err := time.LoadLocation(tz)
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid timezone: %w", err)
	}

	localTime, err := time.Parse("15:04", input)
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid time input: %w", err)
	}

	now := time.Now().In(location)

	localTimeWithDate := time.Date(now.Year(), now.Month(), now.Day(), localTime.Hour(), localTime.Minute(), 0, 0, location)
	return localTimeWithDate.UTC(), nil
}

func ConvertLocalTimeToTime(input string) (time.Time, error) {
	now := time.Now()
	parsed, err := time.Parse("15:04", input)
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid time input: %w", err)
	}
	localTime := time.Date(now.Year(), now.Month(), now.Day(), parsed.Hour(), parsed.Minute(), 0, 0, time.UTC)
	return localTime, nil
}
