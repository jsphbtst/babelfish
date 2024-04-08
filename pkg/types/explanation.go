package types

import "time"

type BreakdownRecord struct {
	Phrase    string    `json:"phrase"`
	Breakdown string    `json:"breakdown"`
	CreatedAt time.Time `json:"createdAt"`
}

type BreakdownJson struct {
	Data []BreakdownRecord `json:"data"`
}
