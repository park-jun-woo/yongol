//ff:func feature=gen-ir type=generator control=sequence
//ff:what buildCSRFConfig -- prepared.Auth → CSRFConfig 변환

package ir

import (
	"github.com/park-jun-woo/yongol/pkg/generate/prepared"
	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

// buildCSRFConfig returns nil when CSRF is not required. Cookie/header
// names resolve through manifest.CsrfConfig.ResolvedCookieName /
// ResolvedHeaderName — the single source of the CSRF defaults.
func buildCSRFConfig(ps *prepared.State) *CSRFConfig {
	if !csrfIsActive(ps) {
		return nil
	}
	var c *manifest.CsrfConfig
	if ps.Auth.Raw != nil {
		c = ps.Auth.Raw.Csrf
	}
	cfg := &CSRFConfig{
		CookieName:       c.ResolvedCookieName(),
		HeaderName:       c.ResolvedHeaderName(),
		MaxAge:           86400,
		Secure:           true,
		HybridBearerSkip: ps.Auth.Mode == "hybrid",
	}
	if c != nil {
		cfg.ExemptPaths = c.ExemptPaths
		if c.MaxAge > 0 {
			cfg.MaxAge = c.MaxAge
		}
	}
	return cfg
}
