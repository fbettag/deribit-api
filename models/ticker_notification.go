package models

type TickerNotification struct {
	Timestamp int64 `json:"timestamp"`
	Stats     struct {
		Volume float64 `json:"volume"`
		Low    float64 `json:"low"`
		High   float64 `json:"high"`
	} `json:"stats"`
	Greeks     struct {
		Delta  float64 `json:"delta"`
		Gamma  float64 `json:"gamma"`
		Rho    float64 `json:"rho"`
		Theta  float64 `json:"theta"`
		Vega   float64 `json:"vega"`
	} `json:"greeks"`
	State           string  `json:"state"`
	SettlementPrice float64 `json:"settlement_price"`
	OpenInterest    float64 `json:"open_interest"`
	MinPrice        float64 `json:"min_price"`
	MaxPrice        float64 `json:"max_price"`
	MarkPrice       float64 `json:"mark_price"`
	AskIV           float64 `json:"ask_iv"`
	BidIV           float64 `json:"bid_iv"`
	InterestRate    float64 `json:"interest_rate"`
	LastPrice       float64 `json:"last_price"`
	InstrumentName  string  `json:"instrument_name"`
	IndexPrice      float64 `json:"index_price"`
	Funding8H       float64 `json:"funding_8h"`
	CurrentFunding  float64 `json:"current_funding"`
	BestBidPrice    float64 `json:"best_bid_price"`
	BestBidAmount   float64 `json:"best_bid_amount"`
	BestAskPrice    float64 `json:"best_ask_price"`
	BestAskAmount   float64 `json:"best_ask_amount"`
}
