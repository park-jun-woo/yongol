//ff:func feature=gen-gogin type=generator control=sequence
//ff:what blockGinRun — http.Server + SIGTERM graceful shutdown + slog + MaxHeaderBytes

package boot

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/yongol"
	"github.com/park-jun-woo/yongol/pkg/generate/gogin/middleware"
)

// defaultHeaderLimit mirrors Go stdlib http.DefaultMaxHeaderBytes (1 MiB).
// Used when manifest.backend.http.header_limit is unset or unparseable.
const defaultHeaderLimit = int64(1 << 20)

// blockGinRun produces the final http.Server run with SIGINT/SIGTERM graceful
// shutdown. Delegates the entire lifecycle (listen, signal wait, shutdown) to
// a sibling top-level helper runServerWithGracefulShutdown so main() stays
// short and carries no depth-1 selection/iteration constructs.
//
// Resolves http.Server.MaxHeaderBytes from manifest.backend.http.header_limit
// (size string, e.g. "1MiB"). Runtime env BACKEND_HTTP_HEADER_LIMIT overrides
// the manifest value; parse failure of the env string falls back to the
// manifest-derived default via envInt64 (matching the BODY_LIMIT pattern).
func blockGinRun(fs *yongol.Fullstack) MainBlock {
	headerLimit := resolveHeaderLimit(fs)
	return MainBlock{
		Name: "gin-run",
		Imports: []string{
			`"context"`,
			`"net/http"`,
			`"os/signal"`,
			`"syscall"`,
			`"time"`,
		},
		Lines: []string{
			fmt.Sprintf(`headerLimit := envInt64("BACKEND_HTTP_HEADER_LIMIT", %d)`, headerLimit),
			`runServerWithGracefulShutdown(r, cancelBootstrap, int(headerLimit))`,
		},
		Funcs: []string{
			ginRunHelperServerLifecycle(),
		},
	}
}

// resolveHeaderLimit returns the manifest-derived MaxHeaderBytes default.
// Nil manifest / empty value / parse failure → 1 MiB (Go stdlib default).
func resolveHeaderLimit(fs *yongol.Fullstack) int64 {
	if fs == nil || fs.Manifest == nil || fs.Manifest.Backend.HTTP == nil {
		return defaultHeaderLimit
	}
	raw := fs.Manifest.Backend.HTTP.HeaderLimit
	if raw == "" {
		return defaultHeaderLimit
	}
	n, err := middleware.ParseSize(raw)
	if err != nil {
		return defaultHeaderLimit
	}
	return n
}
