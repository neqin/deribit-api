package models

type TickerStats struct {
	Volume float64 `json:"volume"`
	Low    float64 `json:"low"`
	High   float64 `json:"high"`
}

type TickerResponse struct {
	BestAskAmount   float64     `json:"best_ask_amount"`
	BestAskPrice    float64     `json:"best_ask_price"`
	BestBidAmount   float64     `json:"best_bid_amount"`
	BestBidPrice    float64     `json:"best_bid_price"`
	CurrentFunding  float64     `json:"current_funding"`
	Funding8H       float64     `json:"funding_8h"`
	IndexPrice      float64     `json:"index_price"`
	InstrumentName  string      `json:"instrument_name"`
	LastPrice       float64     `json:"last_price"`
	MarkPrice       float64     `json:"mark_price"`
	MaxPrice        float64     `json:"max_price"`
	MinPrice        float64     `json:"min_price"`
	OpenInterest    float64     `json:"open_interest"`
	SettlementPrice float64     `json:"settlement_price"`
	State           string      `json:"state"`
	Stats           TickerStats `json:"stats"`
	Timestamp       int64       `json:"timestamp"`
}
