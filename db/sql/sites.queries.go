package db

import "context"

type SitesQuerier interface {
	GetSite(ctx context.Context, key string) (GetSiteQueryResponse, error)
	GenerateSite(ctx context.Context, args GenerateSiteParams) (GenerateSiteQueryResponse, error)
	CountSearch(ctx context.Context, args CountSearchParams) error
	GetSiteStats(ctx context.Context, key string) (GetSiteStatsResponse, error)
	UpdateSite(ctx context.Context, args UpdateSiteParams) (GetSiteQueryResponse, error)
	DeleteSite(ctx context.Context, key string) error
}

var _ SitesQuerier = (*Queries)(nil)
