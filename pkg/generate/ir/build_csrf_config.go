//ff:func feature=gen-ir type=generator control=sequence
//ff:what buildCSRFConfig -- prepared.Auth → CSRFConfig 변환

package ir

import "github.com/park-jun-woo/yongol/pkg/generate/prepared"

// buildCSRFConfig returns nil when CSRF is not required.
func buildCSRFConfig(ps *prepared.State) *CSRFConfig {
	if !csrfIsActive(ps) {
		return nil
	}
	cfg := &CSRFConfig{
		CookieName:       "XSRF-TOKEN",
		HeaderName:       "X-XSRF-TOKEN",
		MaxAge:           86400,
		Secure:           true,
		HybridBearerSkip: ps.Auth.Mode == "hybrid",
	}
	if ps.Auth.Raw != nil && ps.Auth.Raw.Csrf != nil {
		c := ps.Auth.Raw.Csrf
		if c.CookieName != "" {
			cfg.CookieName = c.CookieName
		}
		if c.HeaderName != "" {
			cfg.HeaderName = c.HeaderName
		}
		cfg.ExemptPaths = c.ExemptPaths
		if c.MaxAge > 0 {
			cfg.MaxAge = c.MaxAge
		}
	}
	return cfg
}
