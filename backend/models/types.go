package models

type MeterReading struct {
	Timestamp int64   `json:"timestamp"`
	Kwh       float64 `json:"kwh"`
}

type MarketPrice struct {
	StartTimestamp int64   `json:"start_timestamp"`
	EndTimestamp   int64   `json:"end_timestamp"`
	MarketPrice    float64 `json:"marketprice"`
}
