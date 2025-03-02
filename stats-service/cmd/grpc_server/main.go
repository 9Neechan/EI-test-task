package main

import (
	"context"
	"log"

	"github.com/9Neechan/EI-test-task/stats-service/internal/app"
)

// main is the entry point of the application
func main() {
	ctx := context.Background()

	// Initialize the application
	a, err := app.NewApp(ctx)
	if err != nil {
		log.Fatalf("failed to init app: %s", err.Error())
	}

	// Run the application
	err = a.Run()
	if err != nil {
		log.Fatalf("failed to run app: %s", err.Error())
	}
}

