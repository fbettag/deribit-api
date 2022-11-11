package models

type InstrumentStateNotification struct {
	InstrumentName string `json:"instrument_name"`
	State          string `json:"state"` // created, started, settled, closed, deactivated, terminated
	Timestamp      int64  `json:"timestamp"`
}
