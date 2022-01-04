package main

import (
	"context"
	"fmt"
	"os"

	"github.com/engelmi/gim/internal"
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
		panic(fmt.Sprintf("Error loading config: %+v", err))
	}

	gim, err := internal.NewGopherInTheMiddle(gimconf)
	if err != nil {
		panic(fmt.Sprintf("Error creating Gophers In the Middle: %+v", err))
	}
	gim.Start(ctx, nil)
}

func getRequiredEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		panic(fmt.Sprintf("Required env variable '%s' missing", key))
	}
	return value
}
