package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"

	"foxminded/3.3-weather-forecast-bot/slogger"
)

type Config struct {
	CfgBot     Bot
	CfgHoliday Holiday
	CfgWeather OpenWeather
	CfgGoogle  Google
	CfgMongoDB MongoDB
}

type Bot struct {
	Token   string
	Debug   bool
	Timeout int
}

type Holiday struct {
	Key string
	Url string
}
type OpenWeather struct {
	Key string
	Url string
}

type Google struct {
	Key string
	Url string
}

type MongoDB struct {
	Uri string
}

// LoadConfig loads the configuration for the Telegram bot.
// It retrieves environment variables and optionally loads them from a `.env` file if `useEnvFile` is set to true.
func LoadConfig() Config {

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found — using only system environment")
	}

	var (
		cfg Config
		err error
	)

	//get token
	if cfg.CfgBot.Token = os.Getenv("TELEGRAM_BOT_TOKEN"); cfg.CfgBot.Token == "" {
		log.Fatal("TELEGRAM_BOT_TOKEN env variable not set", err)
	}

	//get debug
	if cfg.CfgBot.Debug, err = strconv.ParseBool(getEnvWithDefault("TELEGRAM_BOT_DEBUG_MODE", "true")); err != nil {
		slogger.Log.Warn("Error parsing TELEGRAM_BOT_DEBUG_MODE", "error", err)
	}
	slogger.Log.Debug("Get debug mode from env", "debug", cfg.CfgBot.Debug)

	//get TELEGRAM_BOT_TIMEOUT
	if cfg.CfgBot.Timeout, err = strconv.Atoi(getEnvWithDefault("TELEGRAM_BOT_TIMEOUT", "5")); err != nil {
		slogger.Log.Warn("Error parsing TELEGRAM_BOT_TIMEOUT", "error", err)
	}
	slogger.Log.Debug("Get TELEGRAM_BOT_TIMEOUT from env", "timeout", cfg.CfgBot.Timeout)

	if cfg.CfgHoliday.Key = os.Getenv("HOLIDAY_API_PRIMARY_KEY"); cfg.CfgHoliday.Key == "" {
		log.Fatal("HOLIDAY_API_PRIMARY_KEY env variable not set", err)
	}

	if cfg.CfgHoliday.Url = os.Getenv("HOLIDAY_API_URL"); cfg.CfgHoliday.Url == "" {
		log.Fatal("HOLIDAY_API_URL env variable not set", err)
	}

	if cfg.CfgWeather.Key = os.Getenv("WEATHER_API_TOKEN"); cfg.CfgWeather.Key == "" {
		log.Fatal("WEATHER_API_TOKEN env variable not set", err)
	}

	if cfg.CfgWeather.Url = os.Getenv("WEATHER_API_URL"); cfg.CfgWeather.Url == "" {
		log.Fatal("WEATHER_API_URL env variable not set", err)
	}

	if cfg.CfgGoogle.Key = os.Getenv("GOOGLE_API_KEY"); cfg.CfgGoogle.Key == "" {
		log.Fatal("GOOGLE_KEY env variable not set", err)
	}
	if cfg.CfgGoogle.Url = os.Getenv("GOOGLE_API_URL"); cfg.CfgGoogle.Url == "" {
		log.Fatal("GOOGLE_URL env variable not set", err)
	}

	if cfg.CfgMongoDB.Uri = os.Getenv("MONGODB_URI"); cfg.CfgMongoDB.Uri == "" {
		log.Fatal("MONGODB_URI env variable not set", err)
	}

	return Config{
		CfgBot: Bot{
			Token:   cfg.CfgBot.Token,
			Debug:   cfg.CfgBot.Debug,
			Timeout: cfg.CfgBot.Timeout,
		},
		CfgHoliday: Holiday{
			Key: cfg.CfgHoliday.Key,
			Url: cfg.CfgHoliday.Url,
		},
		CfgWeather: OpenWeather{
			Key: cfg.CfgWeather.Key,
			Url: cfg.CfgWeather.Url,
		},
		CfgGoogle: Google{
			Key: cfg.CfgGoogle.Key,
			Url: cfg.CfgGoogle.Url,
		},
		CfgMongoDB: MongoDB{
			Uri: cfg.CfgMongoDB.Uri,
		},
	}

}

// getEnvWithDefault get env value or set default
func getEnvWithDefault(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		slogger.Log.Debug("Found environment variable", "key", key, "value", value)
		return value
	}
	slogger.Log.Debug("No environment variable found", "key", key)
	return defaultValue
}
