//ff:type feature=gen-gogin type=generator
//ff:what packageEmitter — one ssac DB-using 패키지의 adapter emit 함수 타입

package infra

import "github.com/park-jun-woo/yongol/pkg/ssacmeta"

// packageEmitter composes the adapter file for one ssac DB-using package.
// Each package requires a purpose-built emitter because the ssac-declared
// Go interface (e.g. cache.CacheModel takes `value any`, returns `(string,
// error)`) does not line up one-to-one with the underlying sqlc query
// signature (e.g. CacheSet takes `value []byte`, CacheGet returns `[]byte`).
//
// interface.yaml remains the single source of truth for *what* the adapter
// must do (which sqlc methods to call, with what params). The emitter's
// glue is limited to type translation at the interface boundary —
//
//   - any → []byte     (JSON marshal on Set-family)
//   - []byte → string  (UTF-8 cast on Get-family)
//   - ttl → expires_at (time.Now().Add on Set-family)
//   - pgx.Tx assert    (queue.PublishTx)
//   - SELECT+UPDATE    (auth.RefreshStore.Consume is not a single query)
type packageEmitter func(iface *ssacmeta.PackageInterface, active []ssacmeta.Port, modulePath, artifactsDir string) error
