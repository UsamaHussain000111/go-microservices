package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/UsamaHussain000111/go-microservices/app"
)

func main() {
	app := app.New()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	err := app.Start(ctx)

	if err != nil {
		fmt.Println("failed to start the app:", err)
	}
}
