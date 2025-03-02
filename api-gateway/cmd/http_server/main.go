package main

import (
	"context"
	"log"

	"github.com/9Neechan/EI-test-task/api-gateway/internal/app"
)

func main() {
	ctx := context.Background()

	// Initialize the application
	a, err := app.NewApp(ctx)
	if err != nil {
		log.Fatalf("failed to initialize the application: %s", err.Error())
	}

	// Run the application
	err = a.Run()
	if err != nil {
		log.Fatalf("failed to start the application: %s", err.Error())
	}
}
