//ff:func feature=gen-ir type=generator control=sequence
//ff:what buildBearerAuthConfig -- prepared.Auth → BearerAuthConfig 변환

package ir

import "github.com/park-jun-woo/yongol/pkg/generate/prepared"

// buildBearerAuthConfig returns nil when auth is not present.
func buildBearerAuthConfig(ps *prepared.State) *BearerAuthConfig {
	if !ps.Auth.Present {
		return nil
	}
	cfg := &BearerAuthConfig{
		Mode: ps.Auth.Mode,
	}
	if ps.Auth.Raw != nil {
		cfg.SecretEnv = ps.Auth.Raw.SecretEnv
		cfg.HasClaims = len(ps.Auth.Raw.Claims) > 0
	}
	if cfg.SecretEnv == "" {
		cfg.SecretEnv = "JWT_SECRET"
	}
	return cfg
}
