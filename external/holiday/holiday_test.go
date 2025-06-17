package holiday

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"foxminded/3.3-weather-forecast-bot/external/holiday/holidayModels"
	"foxminded/3.3-weather-forecast-bot/slogger"
)

type roundTripFunc func(req *http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}

func newTestClient(fn roundTripFunc) *http.Client {
	return &http.Client{
		Transport: fn,
	}
}

const (
	testCase   string = "Test holiday"
	loggerMode bool   = false
)

func TestGetHolidayToday(t *testing.T) {
	slogger.MakeLogger(loggerMode)
	ctx := context.Background()

	tests := []struct {
		name      string
		mockResp  string
		expectMsg string
		mockCode  int
	}{
		{
			name:      "success",
			mockResp:  fmt.Sprintf(`[{"name": "%s"}]`, testCase),
			expectMsg: testCase,
			mockCode:  200,
		},
		{
			name:      "No holidays",
			mockResp:  fmt.Sprintf(`[]`),
			expectMsg: holidayModels.NoHolidays,
			mockCode:  200,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			client := newTestClient(func(req *http.Request) (*http.Response, error) {
				return &http.Response{
						StatusCode: tt.mockCode,
						Body:       io.NopCloser(strings.NewReader(tt.mockResp)),
						Header:     make(http.Header)},
					nil
			})
			service := NewService("", "", client)
			holiday, _ := service.GetHolidayToday(ctx, "US")

			assert.Equal(t, tt.expectMsg, holiday)
		})
	}

}

func TestGetHolidayTodayError(t *testing.T) {
	slogger.MakeLogger(loggerMode)
	ctx := context.Background()

	tests := []struct {
		name      string
		mockResp  string
		expectMsg string
		mockCode  int
	}{
		{
			name:      "Bad status code",
			mockResp:  fmt.Sprintf(`error`),
			mockCode:  500,
			expectMsg: "unexpected response status",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			client := newTestClient(func(req *http.Request) (*http.Response, error) {
				return &http.Response{
						StatusCode: tt.mockCode,
						Body:       io.NopCloser(strings.NewReader(tt.mockResp)),
						Header:     make(http.Header)},
					nil
			})
			service := NewService("", "", client)
			_, err := service.GetHolidayToday(ctx, "US")

			assert.Error(t, err)
			assert.Contains(t, err.Error(), tt.expectMsg)
		})
	}

}
