package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// FlowData model
type FlowData struct {
	gorm.Model
	Timestamp  time.Time `json:"timestamp"`
	FlowRate   float64   `json:"flow_rate"`
	MeterID    string    `json:"meter_id"`
	CustomerID uint      `json:"customer_id"`
}
