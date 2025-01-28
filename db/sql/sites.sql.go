package db

import (
	"context"
	"time"
)

const getSiteQuery = `
	SELECT s.id, s.key, s.long_url, s.created_at, s.updated_at FROM url_shortener_schema.sites s
	WHERE s.key = $1 LIMIT 1
`

func (q *Queries) GetSite(ctx context.Context, key string) (GetSiteQueryResponse, error) {
	row := q.db.QueryRowContext(ctx, getSiteQuery, key)
	var response GetSiteQueryResponse
	err := row.Scan(&response.ID, &response.Key, &response.LongURL, &response.CreatedAt, &response.UpdatedAt)
	return response, err
}

const countSearchQuery = `
	INSERT INTO url_shortener_schema.searches (ip_address, browser, site_id)
	VALUES ($1, $2, $3)
`

type CountSearchParams struct {
	IpAddress string `json:"ip_address"`
	Browser   string `json:"browser"`
	SiteID    int    `json:"site_id"`
}

func (q *Queries) CountSearch(ctx context.Context, args CountSearchParams) error {
	_, err := q.db.ExecContext(ctx, countSearchQuery, args.IpAddress, args.Browser, args.SiteID)
	return err
}

const generateSiteQuery = `
	INSERT INTO url_shortener_schema.sites (key, long_url, salt, expiration)
	VALUES ($1, $2, $3, $4)
	RETURNING id, long_url, created_at, updated_at
`

type GenerateSiteParams struct {
	KEY        string    `json:"key"`
	LONG_URL   string    `json:"long_url"`
	SALT       int64     `json:"salt"`
	EXPIRATION time.Time `json:"expiration"`
}

func (q *Queries) GenerateSite(ctx context.Context, args GenerateSiteParams) (GenerateSiteQueryResponse, error) {
	row := q.db.QueryRowContext(ctx, generateSiteQuery, args.KEY, args.LONG_URL, args.SALT, args.EXPIRATION)
	var response GenerateSiteQueryResponse
	err := row.Scan(&response.ID, &response.LongURL, &response.CreatedAt, &response.UpdatedAt)
	return response, err
}
