package types

import "time"

type HistoryRecord struct {
	Phrase      string    `json:"phrase"`
	Translation string    `json:"translation"`
	CreatedAt   time.Time `json:"createdAt"`
}

type HistoryJson struct {
	Data []HistoryRecord `json:"data"`
}
