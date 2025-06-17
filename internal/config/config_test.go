package config

import (
	"os"
	"testing"

	"foxminded/3.3-weather-forecast-bot/slogger"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	// Инициализация логгера
	slogger.MakeLogger(false)

	// Установка переменных окружения
	if err := os.Setenv("TELEGRAM_BOT_TOKEN", "test_token"); err != nil {
		t.Error(err)
	}
	if err := os.Setenv("TELEGRAM_BOT_DEBUG_MODE", "true"); err != nil {
		t.Error(err)
	}
	if err := os.Setenv("TELEGRAM_BOT_TIMEOUT", "30"); err != nil {
		t.Error(err)
	}
	if err := os.Setenv("HOLIDAY_API_PRIMARY_KEY", "key"); err != nil {
		t.Error(err)
	}
	if err := os.Setenv("HOLIDAY_API_URL", "url"); err != nil {
		t.Error(err)
	}
	if err := os.Setenv("WEATHER_API_TOKEN", "key"); err != nil {
		t.Error(err)
	}
	if err := os.Setenv("WEATHER_API_URL", "url"); err != nil {
		t.Error(err)
	}
	if err := os.Setenv("GOOGLE_API_KEY", "key"); err != nil {
		t.Error(err)
	}
	if err := os.Setenv("GOOGLE_API_URL", "url"); err != nil {
		t.Error(err)
	}
	if err := os.Setenv("MONGODB_URI", "uri"); err != nil {
		t.Error(err)
	}

	value, _ := os.LookupEnv("TELEGRAM_BOT_TIMEOUT")
	slogger.Log.Debug("timeout after setenv", "TELEGRAM_BOT_TIMEOUT", value)

	// Загрузка конфигурации
	cfg := LoadConfig()

	// Проверка результатов
	assert.Equal(t, "test_token", cfg.CfgBot.Token)
	assert.True(t, cfg.CfgBot.Debug)
	assert.Equal(t, 30, cfg.CfgBot.Timeout)
	assert.Equal(t, "key", cfg.CfgHoliday.Key)
	assert.Equal(t, "url", cfg.CfgHoliday.Url)
	assert.Equal(t, "key", cfg.CfgWeather.Key)
	assert.Equal(t, "url", cfg.CfgWeather.Url)
	assert.Equal(t, "key", cfg.CfgGoogle.Key)
	assert.Equal(t, "url", cfg.CfgGoogle.Url)
	assert.Equal(t, "uri", cfg.CfgMongoDB.Uri)

}
