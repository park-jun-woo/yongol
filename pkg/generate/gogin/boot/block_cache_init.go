//ff:func feature=gen-gogin type=generator control=sequence
//ff:what blockCacheInit — cache.Init (postgres 또는 memory) 블록

package boot

import "github.com/park-jun-woo/yongol/pkg/yongol"

// blockCacheInit produces cache initialization. Active when
// manifest.cache.backend is set.
func blockCacheInit(fs *yongol.Fullstack) MainBlock {
	backend := fs.Manifest.Cache.Backend
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
		Name:    "cache-init",
		Active:  hasCache,
		Imports: []string{`"github.com/park-jun-woo/ssac/pkg/cache"`},
		Lines:   lines,
	}
}
