package infra

// cacheWrapperTemplate is the printf-style template for
// arts/backend/internal/infra/cache/postgres.go.
// Header `//ff:` lines are assembled from split string literals so this
// file does not appear to hold two top-of-file annotations.
var cacheWrapperTemplate = cacheWrapperHeaderType + cacheWrapperHeaderWhat + `

package cache

import (
	"context"
	"encoding/json"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/park-jun-woo/ssac/pkg/cache"

	"%[1]s/internal/db"
)

// postgresCache adapts the user's sqlc Queries onto cache.CacheModel.
// The zero value is unusable; construct via NewPostgres(queries).
type postgresCache struct {
	q *db.Queries
}

// NewPostgres returns a cache.CacheModel backed by the user's sqlc Queries.
// Wired from main.go via cache.Init(infracache.NewPostgres(queries)).
func NewPostgres(q *db.Queries) cache.CacheModel {
	return &postgresCache{q: q}
}

// Set marshals value to JSON and upserts a row with expires_at = now+ttl.
// interface.yaml port: %[2]s.
func (c *postgresCache) Set(ctx context.Context, key string, value any, ttl time.Duration) error {
	payload, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return c.q.%[2]s(ctx, db.%[2]sParams{
		Key:       key,
		Value:     payload,
		ExpiresAt: pgtype.Timestamptz{Time: time.Now().Add(ttl), Valid: true},
	})
}

// Get forwards to the sqlc query and returns the raw bytes as a string.
// interface.yaml port: %[3]s.
func (c *postgresCache) Get(ctx context.Context, key string) (string, error) {
	raw, err := c.q.%[3]s(ctx, key)
	if err != nil {
		return "", err
	}
	return string(raw), nil
}

// Delete forwards to the sqlc query. interface.yaml port: %[4]s.
func (c *postgresCache) Delete(ctx context.Context, key string) error {
	return c.q.%[4]s(ctx, key)
}
`

var cacheWrapperHeaderType = "//" + "ff:type feature=infra type=model topic=cache\n"
var cacheWrapperHeaderWhat = "//" + "ff:what postgresCache — ssac/pkg/cache.CacheModel 구현 (yongol codegen from ssac/pkg/cache/interface.yaml)"
