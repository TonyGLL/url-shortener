package db

import (
	"context"
	"time"
)

const getSiteQuery = `
	SELECT s.long_url AS site FROM url_shortener_schema.sites s
	WHERE s.key = $1 LIMIT 1
`

func (q *Queries) GetSite(ctx context.Context, key string) (string, error) {
	row := q.db.QueryRowContext(ctx, getSiteQuery, key)
	var site string
	err := row.Scan(&site)
	return site, err
}

const generateSiteQuery = `
	INSERT INTO url_shortener_schema.sites (key, long_url, salt, expiration)
	VALUES ($1, $2, $3, $4)
`

type GenerateSiteParams struct {
	KEY        string    `json:"key"`
	LONG_URL   string    `json:"long_url"`
	SALT       int64     `json:"salt"`
	EXPIRATION time.Time `json:"expiration"`
}

func (q *Queries) GenerateSite(ctx context.Context, args GenerateSiteParams) error {
	_, err := q.db.ExecContext(ctx, generateSiteQuery, args.KEY, args.LONG_URL, args.SALT, args.EXPIRATION)
	return err
}
