package handlers

import (
	"testing"
)

// TestRealAPIEndpoint ensures the real Awattar API fetches data properly
func TestRealAPIEndpoint(t *testing.T) {
	// Define the start and end timestamps for the request
	startTime := int64(1725400800000)
	endTime := int64(1725408000000)

	prices, err := fetchMarketPrices(startTime, endTime)
	if err != nil {
		t.Fatalf("Failed to fetch market prices from the real API: %v", err)
	}

	// Check if we got the expected number of entries
	if len(prices) != 2 {
		t.Error("Expected 2 records of market data")
	}

	expectedPrices := []float64{108.17, 101.02}
	for i, price := range prices {
		// Ensure the data has been correctly parsed into the MarketPrice type
		if price.StartTimestamp == 0 || price.EndTimestamp == 0 || price.MarketPrice == 0 {
			t.Errorf("Fetched market price has invalid data: %+v", price)
		}
		// Check if we got expected market prices
		if price.MarketPrice != expectedPrices[i] {
			t.Errorf("Fetched market price is not as expected: %.2f", price.MarketPrice)
		}
	}

	// Debug
	//t.Logf("Fetched market prices: %+v", prices)
}
