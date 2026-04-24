//ff:func feature=gen-gogin type=generator control=sequence
//ff:what blockCacheInit — cache.Init (postgres infra 어댑터 또는 memory) 블록

package boot

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/generate/prepared"
)

// blockCacheInit produces cache initialization from a resolved Cache.
// Callers guard with state.ActiveBackends.Cache != nil so this function
// never sees an inactive subsystem — no raw manifest deref possible by
// signature.
//
// Phase002 (ssac/purify) — the "postgres" branch no longer calls
// ssac's cache.NewPostgresCache (removed from ssac in Phase001). It now
// instantiates the yongol-generated adapter at
// `<module>/internal/infra/cache` via cache.NewPostgres(queries), which
// forwards to the user's sqlc CacheSet/Get/Delete queries declared in
// ssac/pkg/cache/interface.yaml.
func blockCacheInit(c prepared.Cache, modulePath string) MainBlock {
	backend := c.Backend
	var lines []string
	imports := []string{
		`"github.com/park-jun-woo/ssac/pkg/cache"`,
	}
	if backend == "postgres" {
		lines = []string{
			`slog.Info("initializing cache (postgres)")`,
			`cache.Init(infracache.NewPostgres(queries))`,
		}
		imports = append(imports, fmt.Sprintf(`infracache "%s/internal/infra/cache"`, modulePath))
	} else {
		lines = []string{
			`slog.Info("initializing cache (memory)")`,
			`cache.Init(cache.NewMemoryCache())`,
		}
	}
	return MainBlock{
		Name: "cache-init",
		// Active left nil: collectActiveBlocks appends this block only
		// when prepared.State.ActiveBackends.Cache != nil.
		Imports: imports,
		Lines:   lines,
	}
}
