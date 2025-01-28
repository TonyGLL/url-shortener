package db

import "time"

type GenerateSiteQueryResponse struct {
	ID        int       `json:"id"`
	LongURL   string    `json:"long_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type GetSiteQueryResponse struct {
	ID        int       `json:"id"`
	Key       string    `json:"key"`
	LongURL   string    `json:"long_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
