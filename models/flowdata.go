package models

import "time"

type FlowData struct {
	ID         uint      `gorm:"primary_key" json:"id"`
	Timestamp  time.Time `json:"timestamp"`
	FlowRate   float64   `json:"flow_rate"`
	MeterID    string    `json:"meter_id"`
	CustomerID uint      `json:"customer_id"`
}
