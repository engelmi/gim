package main

import (
	"context"
	"fmt"

	"github.com/engelmi/gim/internal"
	"github.com/engelmi/gim/pkg/config"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	gimconf, err := config.FromFile("sample.json")
	if err != nil {
		panic(fmt.Sprintf("Error: %+v", err))
	}

	gim, err := internal.NewGopherInTheMiddle(gimconf)
	if err != nil {
		panic(fmt.Sprintf("Error: %+v", err))
	}
	gim.Start(ctx, nil)
}
