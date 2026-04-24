//ff:func feature=gen-gogin type=generator control=sequence topic=cache
//ff:what emitCacheWrapper — ssac/pkg/cache.CacheModel 을 사용자 sqlc Queries 로 구현하는 adapter 를 emit

package infra

import (
	"bytes"
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/ssacmeta"
)

// emitCacheWrapper writes `arts/backend/internal/infra/cache/postgres.go` —
// a postgres adapter that satisfies ssac/pkg/cache.CacheModel by forwarding
// Set/Get/Delete onto the user's sqlc Queries whose names are declared in
// ssac/pkg/cache/interface.yaml (CacheSet / CacheGet / CacheDelete).
//
// Type-translation glue (not catalog):
//
//   - CacheModel.Set takes `value any` with a `time.Duration` ttl; sqlc
//     CacheSet expects `value []byte` with a concrete `expires_at time.Time`.
//     The wrapper marshals value via json.Marshal and computes
//     expires_at = time.Now().Add(ttl), then materialises
//     pgtype.Timestamptz{Time: _, Valid: true} because the generated sqlc
//     params use pgx/v5's pgtype for TIMESTAMPTZ columns.
//
//   - CacheModel.Get returns (string, error); sqlc CacheGet returns
//     ([]byte, error). The wrapper performs a direct []byte → string
//     conversion — callers on the SSaC side may json.Unmarshal the string
//     back into their original type.
//
// Query names are read from the interface.yaml ports verbatim, so renaming
// a port (e.g. CacheSet → CacheUpsert) only requires an interface.yaml
// edit; the wrapper rediscovers the name via portByName.
func emitCacheWrapper(iface *ssacmeta.PackageInterface, active []ssacmeta.Port, modulePath, artifactsDir string) error {
	setPort := portByName(active, "CacheSet")
	getPort := portByName(active, "CacheGet")
	delPort := portByName(active, "CacheDelete")
	if setPort == nil || getPort == nil || delPort == nil {
		// Partial activation — fall back to no emission. The adapter would
		// otherwise fail to satisfy CacheModel, breaking main.go wiring.
		return fmt.Errorf("cache: interface.yaml missing one of CacheSet/CacheGet/CacheDelete (active ports: %d)", len(active))
	}

	var buf bytes.Buffer
	fmt.Fprintf(&buf, `//ff:type feature=infra type=model topic=cache
//ff:what postgresCache — ssac/pkg/cache.CacheModel 구현 (yongol codegen from ssac/pkg/cache/interface.yaml)

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
`, modulePath, setPort.Name, getPort.Name, delPort.Name)

	return writeAdapterFile(artifactsDir, iface.Package, buf.Bytes())
}
