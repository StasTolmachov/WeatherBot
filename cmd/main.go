package main

import (
	"context"
	"flag"
	"os"

	"foxminded/3.3-weather-forecast-bot/internal/config"
	"foxminded/3.3-weather-forecast-bot/internal/db"
	"foxminded/3.3-weather-forecast-bot/internal/server"
	"foxminded/3.3-weather-forecast-bot/internal/worker"
	"foxminded/3.3-weather-forecast-bot/slogger"
)

func main() {
	var logDebugFlag = flag.Bool("logDebug", true, "LogDebug flag")
	flag.Parse()

	slogger.MakeLogger(*logDebugFlag)

	ctx := context.Background()

	cfg := config.LoadConfig()

	mongoDB, err := db.NewMongoDB(ctx, cfg.CfgMongoDB.Uri)
	if err != nil {
		slogger.Log.ErrorContext(ctx, "Failed to connect to MongoDB", "err", err)
		os.Exit(1)
	}
	defer mongoDB.Close(ctx)

	mode := os.Getenv("APP_MODE")
	switch mode {
	case "bot":
		bot := server.NewBot(cfg, mongoDB)

		if err = bot.RunBot(ctx); err != nil {
			slogger.Log.ErrorContext(ctx, "Failed to start bot", "err", err)
			os.Exit(1)
		}
	case "worker":
		workerService := worker.NewService(cfg, mongoDB)

		slogger.Log.InfoContext(ctx, "Weather forecast worker started")
		workerService.Run(ctx)

	}

}
