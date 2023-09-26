package main

import (
	"context"
	"log"

	_ "github.com/sfomuseum/go-libraryofcongress-database-elasticsearch"
	"github.com/sfomuseum/go-libraryofcongress-database/app/query"
)

func main() {

	ctx := context.Background()
	logger := log.Default()

	err := query.Run(ctx, logger)

	if err != nil {
		logger.Fatalf("Failed to run query, %v", err)
	}
}
