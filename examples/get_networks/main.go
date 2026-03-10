package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	admanager "github.com/jfrconley/google-admanager-api-go"
	v202505 "github.com/jfrconley/google-admanager-api-go/services/v202505"
	"github.com/jfrconley/google-admanager-api-go/services/v202505/line_item_service" // types only
)

func main() {
	ctx := context.Background()

	keyFile := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")
	networkCode := os.Getenv("AD_MANAGER_NETWORK_CODE")

	if keyFile == "" || networkCode == "" {
		log.Fatal("Set GOOGLE_APPLICATION_CREDENTIALS and AD_MANAGER_NETWORK_CODE")
	}

	ts, err := admanager.ServiceAccountTokenSourceFromFile(ctx, keyFile)
	if err != nil {
		log.Fatalf("Failed to create token source: %v", err)
	}

	client := admanager.NewClient(ctx, admanager.Config{
		NetworkCode:     networkCode,
		ApplicationName: "admanager-go-example",
	}, ts)

	lineItemSvc := v202505.NewLineItemService(client)

	resp, err := lineItemSvc.GetLineItemsByStatement(&line_item_service.GetLineItemsByStatement{
		FilterStatement: &line_item_service.Statement{
			Query: "WHERE status = 'DELIVERING' LIMIT 10",
		},
	})

	if err != nil {
		log.Fatalf("GetLineItemsByStatement failed: %v", err)
	}

	log.Printf("Line Items: %v\n", resp.Rval.TotalResultSetSize)

	jsonData, err := json.Marshal(resp.Rval)
	if err != nil {
		log.Fatalf("Failed to marshal response: %v", err)
	}

	log.Printf("Line Items: %v\n", string(jsonData))
}
