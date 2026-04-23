//ff:func feature=gen-gogin type=util control=sequence
//ff:what resolveAuthInitConfig — prepared.Auth 에서 blockAuthInit 가 필요한 모든 값 해석

package boot

import "github.com/park-jun-woo/yongol/pkg/generate/prepared"

// resolveAuthInitConfig walks prepared.Auth (and its underlying Cookie
// sub-block) returning a fully populated authInitConfig with defaults
// applied. Mode is read from prepared.Auth.Mode which is always the
// defaulted value — no raw ResolvedMode() call remains.
func resolveAuthInitConfig(a prepared.Auth) authInitConfig {
	cfg := authInitConfig{
		SecretEnv:   "JWT_SECRET",
		AccessTTL:   "15m",
		RefreshTTL:  "168h",
		Mode:        "cookie",
		SameSite:    "Lax",
		AccessName:  "__Host-access_token",
		RefreshName: "__Host-refresh_token",
	}
	if !a.Present || a.Raw == nil {
		return cfg
	}
	raw := a.Raw
	if raw.SecretEnv != "" {
		cfg.SecretEnv = raw.SecretEnv
	}
	if raw.AccessTokenTTL != "" {
		cfg.AccessTTL = raw.AccessTokenTTL
	}
	if raw.RefreshTokenTTL != "" {
		cfg.RefreshTTL = raw.RefreshTokenTTL
	}
	cfg.Mode = a.Mode
	cfg.DetectReuse = raw.DetectReuseLogoutAll
	resolveAuthCookieConfig(raw.Cookie, &cfg)
	return cfg
}
