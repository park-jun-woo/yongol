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
	fmt.Fprintf(&buf, cacheWrapperTemplate, modulePath, setPort.Name, getPort.Name, delPort.Name)

	return writeAdapterFile(artifactsDir, iface.Package, buf.Bytes())
}
