//ff:func feature=gen-gogin type=generator control=sequence topic=observability
//ff:what blockPrometheus — middleware.PrometheusMiddleware + /metrics 라우팅 등록

package boot

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// blockPrometheus emits the Prometheus middleware registration and mounts
// the scrape endpoint. BACKEND_OBSERVABILITY_METRICS_ENABLED (env override)
// lets operators disable the middleware + route at runtime without
// regenerating. BACKEND_OBSERVABILITY_METRICS_PATH overrides the route path
// (still required to start with "/" — guarded by OBS-001 at validate-time).
//
// Ordering: registered AFTER blockRouter so r (gin.Engine) is in scope, and
// BEFORE blockRegisterHandlers so every oapi-codegen route inherits the
// middleware. /metrics itself is mounted directly on r to bypass auth/authz
// middlewares that attach later in the chain.
func blockPrometheus(fs *yongol.Fullstack, modulePath string) MainBlock {
	if !hasPrometheus(fs) {
		return MainBlock{Name: "prometheus"}
	}
	path := prometheusPath(fs)

	lines := []string{
		`promEnabled := envBool("BACKEND_OBSERVABILITY_METRICS_ENABLED", true)`,
		fmt.Sprintf(`promPath := envString("BACKEND_OBSERVABILITY_METRICS_PATH", %q)`, path),
		`if promEnabled {`,
		`	r.Use(middleware.PrometheusMiddleware())`,
		`	r.GET(promPath, middleware.PrometheusHandler())`,
		`}`,
	}

	imports := []string{
		fmt.Sprintf(`"%s/internal/middleware"`, modulePath),
	}

	return MainBlock{
		Name:    "prometheus",
		Imports: imports,
		Lines:   lines,
	}
}
