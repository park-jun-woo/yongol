//ff:func feature=gen-gogin type=generator control=sequence topic=request-id
//ff:what blockRequestID — middleware.RequestID(...) 최상위 등록 (Phase004)

package boot

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// blockRequestID emits r.Use(middleware.RequestID(...)) at the very top of
// the gin middleware chain so every downstream middleware (CORS, prometheus,
// body_limit, rate_limit, request_validator, handlers) observes the id via
// gin.Context / request.Context(). Runtime policy comes from
// manifest.backend.error.request_id (trust_upstream, header) with env
// overrides:
//
//	BACKEND_ERROR_REQUEST_ID_TRUST_UPSTREAM
//	BACKEND_ERROR_REQUEST_ID_HEADER
//
// Always active — the id cost is negligible (ULID) and downstream middlewares
// assume the key is always set.
func blockRequestID(fs *yongol.Fullstack, modulePath string) MainBlock {
	trustUpstream, header := resolveRequestIDConfig(fs)
	lines := []string{
		fmt.Sprintf(`ridTrustUpstream := envBool("BACKEND_ERROR_REQUEST_ID_TRUST_UPSTREAM", %v)`, trustUpstream),
		fmt.Sprintf(`ridHeader := envString("BACKEND_ERROR_REQUEST_ID_HEADER", %q)`, header),
		`r.Use(middleware.RequestID(ridTrustUpstream, ridHeader))`,
	}
	return MainBlock{
		Name: "request-id",
		Imports: []string{
			fmt.Sprintf(`"%s/internal/middleware"`, modulePath),
		},
		Lines: lines,
	}
}

// resolveRequestIDConfig reads manifest.backend.error.request_id with
// sensible defaults when the block is absent. trust_upstream defaults to
// true; header defaults to "X-Request-Id".
func resolveRequestIDConfig(fs *yongol.Fullstack) (bool, string) {
	trust := true
	header := "X-Request-Id"
	if fs == nil || fs.Manifest == nil {
		return trust, header
	}
	e := fs.Manifest.Backend.Error
	if e == nil || e.RequestID == nil {
		return trust, header
	}
	if e.RequestID.TrustUpstream != nil {
		trust = *e.RequestID.TrustUpstream
	}
	if e.RequestID.Header != "" {
		header = e.RequestID.Header
	}
	return trust, header
}
