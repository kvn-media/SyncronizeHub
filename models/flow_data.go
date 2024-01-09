// models/flow_data.go

package models

import "time"

// FlowDataModel represents flow data in the system.
type FlowDataModel struct {
	Timestamp time.Time `json:"timestamp"`
	FlowRate  float64   `json:"flow_rate"`
	MeterID   string    `json:"meter_id"`
	CustomerID string   `json:"customer_id"`
}

// NewFlowDataModel creates a new instance of FlowDataModel with the given parameters.
func NewFlowDataModel(timestamp time.Time, flowRate float64, meterID, customerID string) *FlowDataModel {
	return &FlowDataModel{
		Timestamp:  timestamp,
		FlowRate:   flowRate,
		MeterID:    meterID,
		CustomerID: customerID,
	}
}

// Validate method for FlowDataModel to perform validation checks.
func (fd *FlowDataModel) Validate() error {
	// Perform validation checks, e.g., check if timestamp is not in the future, flow rate is valid, etc.
	// Return an error if validation fails.
	return nil
}
