//ff:func feature=gen-gogin type=generator control=iteration dimension=2
//ff:what EmitAll — fs.SsacInterfaces 순회하여 각 DB-using 패키지의 postgres.go 작성

package infra

import (
	"fmt"
	"sort"

	"github.com/park-jun-woo/yongol/pkg/ssacmeta"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

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
//
// Anything else — parameter order, type names, query names — is read from
// the PackageInterface so adding a new port needs only a new entry in
// interface.yaml, not a wrapper edit.
type packageEmitter func(iface *ssacmeta.PackageInterface, active []ssacmeta.Port, modulePath, artifactsDir string) error

// packageEmitters dispatches from `interface.yaml` package name to the
// concrete emitter. Packages that declare no Go-interface mismatch (none at
// time of Phase002 write — every DB-using ssac package currently needs at
// least one glue conversion) could fall back to emitPostgresImplGeneric.
var packageEmitters = map[string]packageEmitter{
	"cache":   emitCacheWrapper,
	"session": emitSessionWrapper,
	"queue":   emitQueueWrapper,
	"auth":    emitAuthWrapper,
}

// EmitAll iterates over fs.SsacInterfaces and writes
// `<artifactsDir>/backend/internal/infra/<pkg>/postgres.go` for every
// DB-using ssac package whose declared ports include at least one active
// port under the current manifest.
//
// Inactive ports (`when:` evaluated false) are filtered; when every port of
// a package is inactive the package emits nothing (no file). This keeps the
// artifacts tree free of unused adapters when e.g. `manifest.cache.backend`
// is "memory".
//
// modulePath is the user project's Go module path and is embedded in the
// generated file's import for `<modulePath>/internal/db`.
func EmitAll(fs *yongol.Fullstack, artifactsDir, modulePath string) error {
	if fs == nil || len(fs.SsacInterfaces) == 0 {
		return nil
	}
	mctx := flattenManifest(fs)

	// Deterministic order.
	pkgs := make([]string, 0, len(fs.SsacInterfaces))
	for k := range fs.SsacInterfaces {
		pkgs = append(pkgs, k)
	}
	sort.Strings(pkgs)

	for _, pkg := range pkgs {
		iface := fs.SsacInterfaces[pkg]
		if iface == nil {
			continue
		}
		// Dynamic-port packages (authz) are handled by Phase003 owner-lookup
		// emitter, not here — no static ports to walk.
		if len(iface.Ports) == 0 {
			continue
		}
		active := activePorts(iface.Ports, mctx)
		if len(active) == 0 {
			continue
		}
		emitter, ok := packageEmitters[pkg]
		if !ok {
			// No dedicated emitter registered — skip. A future package with
			// a pure sqlc-forwarder interface could be wired to
			// emitPostgresImplGeneric here.
			continue
		}
		if err := emitter(iface, active, modulePath, artifactsDir); err != nil {
			return fmt.Errorf("infra: emit %s: %w", pkg, err)
		}
	}
	return nil
}

func activePorts(ports []ssacmeta.Port, mctx map[string]any) []ssacmeta.Port {
	var out []ssacmeta.Port
	for _, p := range ports {
		if ssacmeta.EvaluateWhen(p.When, mctx) {
			out = append(out, p)
		}
	}
	return out
}

// portByName locates a port from the active subset by its interface.yaml
// name (e.g. "CacheSet"). Returns nil when the port is inactive under the
// current manifest — callers must handle the missing port (typically by
// skipping the method that depends on it).
func portByName(active []ssacmeta.Port, name string) *ssacmeta.Port {
	for i := range active {
		if active[i].Name == name {
			return &active[i]
		}
	}
	return nil
}

// flattenManifest builds the map[string]any that ssacmeta.EvaluateWhen
// consults. Paths referenced by interface.yaml `when:` expressions must match
// the shape produced here.
//
// Supported paths today:
//   manifest.cache.backend
//   manifest.session.backend
//   manifest.queue.backend
//   manifest.backend.auth.refresh.enabled
//
// When fs.Manifest is nil we return an empty map — every `when:` on a
// non-"always" port resolves to false so nothing is emitted. This matches the
// expected behavior for memory-only dev runs.
func flattenManifest(fs *yongol.Fullstack) map[string]any {
	m := map[string]any{}
	if fs == nil || fs.Manifest == nil {
		return m
	}
	cfg := fs.Manifest
	// Infrastructure backends (cache/session/queue) live at the top level
	// of manifest.yaml — they are siblings of `backend:`, not children.
	if cfg.Cache != nil {
		m["cache"] = map[string]any{"backend": cfg.Cache.Backend}
	}
	if cfg.Session != nil {
		m["session"] = map[string]any{"backend": cfg.Session.Backend}
	}
	if cfg.Queue != nil {
		m["queue"] = map[string]any{"backend": cfg.Queue.Backend}
	}
	// backend.auth.refresh.enabled — auth package is the only DB-using
	// subsystem under `backend:`, and refresh_tokens support is implicit
	// whenever `backend.auth` is declared (there is no sub-struct today;
	// the ssac auth package either supports refresh rotation or it does
	// not, and the user opts in simply by having auth configured).
	backend := map[string]any{}
	if cfg.Backend.Auth != nil {
		backend["auth"] = map[string]any{
			"refresh": map[string]any{"enabled": true},
		}
	}
	m["backend"] = backend
	return m
}
