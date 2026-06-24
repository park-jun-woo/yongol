//ff:func feature=gen-gogin type=util control=sequence
//ff:what domainWireMode — 도메인 strict 미들웨어 배선용 auth_mode 해석 (override → backend → bearer)

package boot

import (
	"github.com/park-jun-woo/yongol/pkg/generate/prepared"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// domainWireMode resolves the auth transport mode used to pick a domain's
// strict middleware (Phase008). A domain's own auth_mode wins; otherwise it
// inherits the resolved backend.auth.mode (prepared.AuthFor applies the
// jwt→bearer and empty→cookie defaults). When no backend auth block exists the
// route group is still wired for the opt-in bearerAuth middleware, so the
// fallback is "bearer" — matching the manifest middleware that gated it.
func domainWireMode(fs *yongol.Fullstack, name string) string {
	cfg := fs.Manifest.Domains[name]
	if cfg.AuthMode != "" {
		return cfg.AuthMode
	}
	if fs.Manifest.Backend.Auth != nil {
		if m := prepared.AuthFor(fs).Mode; m != "" {
			return m
		}
	}
	return "bearer"
}
