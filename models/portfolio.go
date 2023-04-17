package models

type Portfolio struct {
	AvailableFunds           float64    `json:"available_funds"`
	AvailableWithdrawalFunds float64    `json:"available_withdrawal_funds"`
	Balance                  float64    `json:"balance"`
	Currency                 string `json:"currency"`
	Equity                   float64    `json:"equity"`
	InitialMargin            float64    `json:"initial_margin"`
	MaintenanceMargin        float64    `json:"maintenance_margin"`
	MarginBalance            float64    `json:"margin_balance"`
}
