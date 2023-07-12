package models

import "time"

// URL model
type URL struct {
	LongUrl   string    `json:"longURL"`
	ShortUrl  string    `json:"shortUrl"`
	CreatedAt time.Time `json:"createdAt"`
}
