package utils

import (
	"challenge.zaehlerfreunde.com/models"
	"sort"
)

// SortMeterReadings sorts the MeterReading data by Timestamp.
func SortMeterReadings(readings []models.MeterReading) {
	sort.Slice(readings, func(i, j int) bool {
		return readings[i].Timestamp < readings[j].Timestamp
	})
}
