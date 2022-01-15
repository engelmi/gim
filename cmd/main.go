package main

import (
	"context"
	"os"

	"github.com/engelmi/gim/internal"
	"github.com/engelmi/gim/internal/logger"
	"github.com/engelmi/gim/pkg/config"
	"github.com/engelmi/gim/pkg/contract"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	configJsonStr := getRequiredEnv(contract.ENV_KEY_CONFIG)
	gimconf, err := config.FromJsonString(configJsonStr)
	if err != nil {
		logger.GetLogger().WithError(err).Fatal("Error loading config")
	}
	logger.SetLogLevel(gimconf.Logger.Level)

	l := logger.GetLogger()
	l.Info("Setting up gophers...")
	gim, err := internal.NewGopherInTheMiddle(gimconf)
	if err != nil {
		logger.GetLogger().WithError(err).Fatal("Error creating Gophers In the Middle")
	}

	gim.Start(ctx, nil)
}

func getRequiredEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		logger.GetLogger().WithField("missing-key", key).Fatal("Required env variable")
	}
	return value
}
