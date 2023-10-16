package main

import (
	"context"
	"fmt"

	"github.com/brad4au67/go-chi-microservice/app"
)

func main() {
	app := app.New()
	err := app.Start(context.TODO())
	if err != nil {
		fmt.Println("Failed to start app:", err)
	}
}
