//ff:func feature=gen-gogin type=util control=sequence
//ff:what resolveAuthInitConfig — manifest.backend.auth 에서 blockAuthInit 가 필요한 모든 값 해석

package boot

import "github.com/park-jun-woo/yongol/pkg/yongol"

// resolveAuthInitConfig walks manifest.backend.auth (and its Cookie sub-
// block) returning a fully populated authInitConfig with defaults applied.
// Flat, selection-free code keeps the helper inside Q1 sequence limits;
// the deep Cookie sub-block is resolved by resolveAuthCookieConfig.
func resolveAuthInitConfig(fs *yongol.Fullstack) authInitConfig {
	cfg := authInitConfig{
		SecretEnv:   "JWT_SECRET",
		AccessTTL:   "15m",
		RefreshTTL:  "168h",
		Mode:        "cookie",
		SameSite:    "Lax",
		AccessName:  "__Host-access_token",
		RefreshName: "__Host-refresh_token",
	}
	if fs.Manifest == nil || fs.Manifest.Backend.Auth == nil {
		return cfg
	}
	a := fs.Manifest.Backend.Auth
	if a.SecretEnv != "" {
		cfg.SecretEnv = a.SecretEnv
	}
	if a.AccessTokenTTL != "" {
		cfg.AccessTTL = a.AccessTokenTTL
	}
	if a.RefreshTokenTTL != "" {
		cfg.RefreshTTL = a.RefreshTokenTTL
	}
	cfg.Mode = a.ResolvedMode()
	cfg.DetectReuse = a.DetectReuseLogoutAll
	resolveAuthCookieConfig(a.Cookie, &cfg)
	return cfg
}
