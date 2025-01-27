package db

import "context"

type SitesQuerier interface {
	GetSite(ctx context.Context, key string) (string, error)
	GenerateSite(ctx context.Context, args GenerateSiteParams) error
}

var _ SitesQuerier = (*Queries)(nil)
