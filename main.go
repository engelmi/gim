package main

import (
	"context"
	"fmt"

	"github.com/engelmi/gim/config"
	"github.com/engelmi/gim/gim"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	gimconf, err := config.FromFile("sample.json")
	if err != nil {
		panic(fmt.Sprintf("Error: %+v", err))
	}

	gim, err := gim.NewGopherInTheMiddle(gimconf)
	if err != nil {
		panic(fmt.Sprintf("Error: %+v", err))
	}
	gim.Start(ctx, nil)
}
