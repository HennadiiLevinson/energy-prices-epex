package handlers

import (
	"challenge.zaehlerfreunde.com/models"
	"challenge.zaehlerfreunde.com/utils"
	"testing"
)

const (
	// Constants for timestamps between 12am-3am (in Unix milliseconds)
	TS_12AM = 1725487200000
	TS_1AM  = 1725490800000
	TS_2AM  = 1725494400000
	TS_3AM  = 1725498000000
	// Constants for meter readings between 12am-3am (in kWh)
	KWH_12AM = 230000.0 // 12am reading
	KWH_1AM  = 230100.0 // 1am reading
	KWH_2AM  = 230200.0 // 2am reading
	KWH_3AM  = 230500.0 // 3am reading
	// Constants for Market Prices (in Euro/MWh)
	MP_12AM_1AM = 120.0 // 12am - 1am price
	MP_1AM_2AM  = 150.0 // 1am - 2am price
	MP_2AM_3AM  = 100.0 // 2am - 3am price
)

// Table-driven test for CalculateEnergyCost
func TestCalculateEnergyCost(t *testing.T) {
	tests := []struct {
		name     string
		readings []models.MeterReading
		prices   []models.MarketPrice
		want     float64
	}{
		{
			name: "Test with normal consumption",
			readings: []models.MeterReading{
				{Timestamp: TS_12AM, Kwh: KWH_12AM},
				{Timestamp: TS_1AM, Kwh: KWH_1AM},
				{Timestamp: TS_2AM, Kwh: KWH_2AM},
				{Timestamp: TS_3AM, Kwh: KWH_3AM},
			},
			prices: []models.MarketPrice{
				{StartTimestamp: TS_12AM, EndTimestamp: TS_1AM, MarketPrice: MP_12AM_1AM},
				{StartTimestamp: TS_1AM, EndTimestamp: TS_2AM, MarketPrice: MP_1AM_2AM},
				{StartTimestamp: TS_2AM, EndTimestamp: TS_3AM, MarketPrice: MP_2AM_3AM},
			},
			want: ((KWH_1AM - KWH_12AM) * MP_12AM_1AM / models.PriceUnitDivider) +
				((KWH_2AM - KWH_1AM) * MP_1AM_2AM / models.PriceUnitDivider) +
				((KWH_3AM - KWH_2AM) * MP_2AM_3AM / models.PriceUnitDivider),
		},
		{
			name: "Test with zero consumption",
			readings: []models.MeterReading{
				{Timestamp: TS_12AM, Kwh: KWH_12AM},
				{Timestamp: TS_1AM, Kwh: KWH_12AM}, // No consumption between these two readings
			},
			prices: []models.MarketPrice{
				{StartTimestamp: TS_12AM, EndTimestamp: TS_1AM, MarketPrice: MP_12AM_1AM},
			},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := calculateEnergyCost(tt.readings, tt.prices)
			if actual != tt.want {
				t.Errorf("got %.2f, want %.2f", actual, tt.want)
			}
			// Debug
			// log.Printf("Energy cost\n Expected:%.2f\n Actual:%.2f\n", tt.want, actual)
		})
	}
}

// Table-driven test for SortMeterReadings
func TestSortMeterReadings(t *testing.T) {
	tests := []struct {
		name     string
		readings []models.MeterReading
		want     []models.MeterReading
	}{
		{
			name: "Test with unsorted meter readings",
			readings: []models.MeterReading{
				{Timestamp: TS_3AM, Kwh: KWH_3AM},
				{Timestamp: TS_12AM, Kwh: KWH_12AM},
				{Timestamp: TS_1AM, Kwh: KWH_1AM},
				{Timestamp: TS_2AM, Kwh: KWH_2AM},
			},
			want: []models.MeterReading{
				{Timestamp: TS_12AM, Kwh: KWH_12AM},
				{Timestamp: TS_1AM, Kwh: KWH_1AM},
				{Timestamp: TS_2AM, Kwh: KWH_2AM},
				{Timestamp: TS_3AM, Kwh: KWH_3AM},
			},
		},
		{
			name: "Test with already sorted meter readings",
			readings: []models.MeterReading{
				{Timestamp: TS_12AM, Kwh: KWH_12AM},
				{Timestamp: TS_1AM, Kwh: KWH_1AM},
				{Timestamp: TS_2AM, Kwh: KWH_2AM},
				{Timestamp: TS_3AM, Kwh: KWH_3AM},
			},
			want: []models.MeterReading{
				{Timestamp: TS_12AM, Kwh: KWH_12AM},
				{Timestamp: TS_1AM, Kwh: KWH_1AM},
				{Timestamp: TS_2AM, Kwh: KWH_2AM},
				{Timestamp: TS_3AM, Kwh: KWH_3AM},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			utils.SortMeterReadings(tt.readings)
			for i := range tt.readings {
				if tt.readings[i] != tt.want[i] {
					t.Errorf("got %+v, want %+v", tt.readings[i], tt.want[i])
				}
			}
		})
	}
}
