//ff:func feature=gen-ir type=generator control=sequence
//ff:what buildAuthInfraConfig -- manifest + prepared.Auth → AuthInfraConfig 변환

package ir

import (
	"github.com/park-jun-woo/yongol/pkg/generate/prepared"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// buildAuthInfraConfig extracts auth infrastructure settings from
// manifest + prepared.State.
func buildAuthInfraConfig(fs *yongol.Fullstack, ps *prepared.State) *AuthInfraConfig {
	cfg := &AuthInfraConfig{
		Mode:            ps.Auth.Mode,
		SecretEnv:       "JWT_SECRET",
		AccessTokenTTL:  "15m",
		RefreshTokenTTL: "168h",
	}
	if ps.Auth.Raw == nil {
		return cfg
	}
	a := ps.Auth.Raw
	if a.SecretEnv != "" {
		cfg.SecretEnv = a.SecretEnv
	}
	if a.AccessTokenTTL != "" {
		cfg.AccessTokenTTL = a.AccessTokenTTL
	}
	if a.RefreshTokenTTL != "" {
		cfg.RefreshTokenTTL = a.RefreshTokenTTL
	}
	_ = fs // reserved for future use (e.g. extracting DDL-driven config)
	return cfg
}
