package main

import (
	"context"
	"log"

	_ "github.com/sfomuseum/go-libraryofcongress-database-elasticsearch"
	"github.com/sfomuseum/go-libraryofcongress-database/app/index"
)

func main() {

	ctx := context.Background()
	logger := log.Default()

	err := index.Run(ctx, logger)

	if err != nil {
		logger.Fatalf("Failed to run indexer, %v", err)
	}
}
