//ff:func feature=gen-gogin type=generator control=sequence
//ff:what blockCacheInit — cache.Init (postgres 또는 memory) 블록

package boot

import "github.com/park-jun-woo/yongol/pkg/generate/prepared"

// blockCacheInit produces cache initialization from a resolved Cache.
// Callers guard with state.ActiveBackends.Cache != nil so this function
// never sees an inactive subsystem — no raw manifest deref possible by
// signature.
func blockCacheInit(c prepared.Cache) MainBlock {
	backend := c.Backend
	var lines []string
	if backend == "postgres" {
		lines = []string{
			`slog.Info("initializing cache (postgres)")`,
			`cm, err := cache.NewPostgresCache(ctx, conn)`,
			`if err != nil {`,
			`	slog.Error("cache init", "err", err)`,
			`	os.Exit(1)`,
			`}`,
			`cache.Init(cm)`,
		}
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
		Imports: []string{`"github.com/park-jun-woo/ssac/pkg/cache"`},
		Lines:   lines,
	}
}
