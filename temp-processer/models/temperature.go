package models

import "time"

type TemperatureReading struct {
	SensorID    string    `json:"sensor_id"`
	Temperature float64   `json:"temperature"`
	Timestamp   time.Time `json:"timestamp,omitempty"`
}
