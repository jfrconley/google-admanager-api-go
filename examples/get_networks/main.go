package main

import (
	"context"
	"log"
	"os"

	admanager "google-admanager-api-go"
	v202505 "google-admanager-api-go/services/v202505"
	"google-admanager-api-go/services/v202505/line_item_service"
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

	//networkSvc := network_service.NewNetworkServiceInterface(
	//	v202505.NewService(client, "NetworkService"),
	//)

	//resp, err := networkSvc.GetCurrentNetwork(&network_service.GetCurrentNetwork{})
	//if err != nil {
	//	log.Fatalf("GetCurrentNetwork failed: %v", err)
	//}
	//
	//n := resp.Rval
	//fmt.Printf("Network: %s (code: %s)\n", n.DisplayName, n.NetworkCode)
	//fmt.Printf("  Currency:  %s\n", n.CurrencyCode)
	//fmt.Printf("  Time Zone: %s\n", n.TimeZone)
	//fmt.Printf("  Test:      %v\n", n.IsTest)

	lineItemSvc := line_item_service.NewLineItemServiceInterface(v202505.NewService(client, "LineItemService"))

	resp, err := lineItemSvc.GetLineItemsByStatement(&line_item_service.GetLineItemsByStatement{
		FilterStatement: &line_item_service.Statement{
			Query: "WHERE status = 'DELIVERING' LIMIT 10",
		},
	})

	if err != nil {
		log.Fatalf("GetLineItemsByStatement failed: %v", err)
	}

	log.Printf("Line Items: %v\n", resp.Rval.TotalResultSetSize)

}
