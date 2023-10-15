package main

import (
	"context"
	"fmt"

	"github.com/UsamaHussain000111/go-microservices/app"
)

func main() {
	app := app.New()

	err := app.Start(context.TODO())

	if err != nil {
		fmt.Println("failed to start the app:", err)
	}
}
