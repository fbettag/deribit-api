package models

type SubmitTransferToSubaccountParams struct {
	Currency string `json:"currency"`
	Amount float64 `json:"amount"`
	Destination int `json:"destination"`
}
