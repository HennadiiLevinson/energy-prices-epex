package handlers

import (
	"challenge.zaehlerfreunde.com/models"
	"challenge.zaehlerfreunde.com/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// EnergyCostHandler is the handler for the /energy_cost endpoint
func EnergyCostHandler(c *gin.Context) {
	var readings []models.MeterReading

	if err := c.ShouldBindJSON(&readings); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Println(err.Error())
		return
	}

	// Ensure there are readings to process, and they have the right length
	if len(readings) == 0 || len(readings) < models.MeterReadingsCount {
		errMsg := "Wrong meter readings data format or none"
		c.JSON(http.StatusBadRequest, gin.H{"error": errMsg})
		log.Println(errMsg)
		return
	}

	// Make sure readings are sorted properly
	utils.SortMeterReadings(readings)

	startTime := readings[0].Timestamp
	endTime := readings[len(readings)-1].Timestamp

	// Fetch EPEX spot prices for the relevant time range
	prices, err := fetchMarketPrices(startTime, endTime)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch market prices"})
		fmt.Println(err.Error())
		return
	}

	totalCost := calculateEnergyCost(readings, prices)

	c.JSON(http.StatusOK, gin.H{"total_cost": totalCost})
}

// Fetch market prices from Awattar API for the given time range (start and end timestamps)
func fetchMarketPrices(start, end int64) ([]models.MarketPrice, error) {
	// Construct the Awattar API URL with start and end time in milliseconds
	url := fmt.Sprintf("%s?start=%d&end=%d", models.AwattarAPIURL, start, end)
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Error fetching market data: %s", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response body: %s", err)
		return nil, err
	}

	var marketData struct {
		Data []models.MarketPrice `json:"data"`
	}

	if err := json.Unmarshal(body, &marketData); err != nil {
		log.Printf("Error unmarshalling market data: %s", err)
		return nil, err
	}

	return marketData.Data, nil
}

// Calculate the energy cost for the provided meter readings and market prices
func calculateEnergyCost(readings []models.MeterReading, prices []models.MarketPrice) float64 {
	// Perform the cost calculation using the same index for both meterReadings and marketPrices
	var totalCost float64

	for i := 1; i < len(readings); i++ {
		consumption := readings[i].Kwh - readings[i-1].Kwh
		pricePerKwh := prices[i-1].MarketPrice / models.PriceUnitDivider // Convert price from Euro/MWh to Euro/kWh
		totalCost += consumption * pricePerKwh
	}
	return totalCost
}
