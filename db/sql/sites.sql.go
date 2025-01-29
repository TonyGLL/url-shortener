package db

import (
	"context"
	"time"
)

/* GET SITE */
const getSiteQuery = `
	SELECT s.id, s.key, s.long_url, s.created_at, s.updated_at FROM url_shortener_schema.sites s
	WHERE s.key = $1 AND s.deleted IS NOT TRUE LIMIT 1
`

func (q *Queries) GetSite(ctx context.Context, key string) (GetSiteQueryResponse, error) {
	row := q.db.QueryRowContext(ctx, getSiteQuery, key)
	var response GetSiteQueryResponse
	err := row.Scan(&response.ID, &response.Key, &response.LongURL, &response.CreatedAt, &response.UpdatedAt)
	return response, err
}

/* GET SITE */

/* COUNT SEARCHES */
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

/* COUNT SEARCHES */

/* GENERATE SITE */
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

/* GENERATE SITE */

/* GET SITE STATS */
const getSiteStatsQuery = `
	SELECT s.id, s.key, s.long_url, s.created_at, s.updated_at, COUNT(se.site_id) AS accessCount
	FROM url_shortener_schema.sites s
	LEFT JOIN url_shortener_schema.searches se ON s.id = se.site_id
	WHERE s.key = $1 AND s.deleted IS NOT TRUE
	GROUP BY s.id, s.long_url;
`

func (q *Queries) GetSiteStats(ctx context.Context, key string) (GetSiteStatsResponse, error) {
	row := q.db.QueryRowContext(ctx, getSiteStatsQuery, key)
	var response GetSiteStatsResponse
	err := row.Scan(&response.ID, &response.Key, &response.LongURL, &response.CreatedAt, &response.UpdatedAt, &response.AccessCount)
	return response, err
}

/* GET SITE STATS */

/* UPDATE SITE */
const updateSiteQuery = `
	UPDATE url_shortener_schema.sites s 
	SET long_url = $2 
	WHERE s."key" = $1
	RETURNING s.id, s."key", s.long_url, s.created_at, s.updated_at;
`

type UpdateSiteParams struct {
	KEY     string `json:"key"`
	LongURL string `json:"long_url"`
}

func (q *Queries) UpdateSite(ctx context.Context, args UpdateSiteParams) (GetSiteQueryResponse, error) {
	row := q.db.QueryRowContext(ctx, updateSiteQuery, args.KEY, args.LongURL)
	var response GetSiteQueryResponse
	err := row.Scan(&response.ID, &response.Key, &response.LongURL, &response.CreatedAt, &response.UpdatedAt)
	return response, err
}

/* UPDATE SITE */

/* DELETE SITE */
const deleteSiteQuery = `
	UPDATE url_shortener_schema.sites s 
	SET deleted = TRUE 
	WHERE s.key = $1
`

func (q *Queries) DeleteSite(ctx context.Context, key string) error {
	_, err := q.db.ExecContext(ctx, deleteSiteQuery, key)
	return err
}

/* DELETE SITE */
