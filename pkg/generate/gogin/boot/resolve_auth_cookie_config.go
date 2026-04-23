//ff:func feature=gen-gogin type=util control=sequence
//ff:what resolveAuthCookieConfig — manifest.auth.cookie 서브블록 값으로 authInitConfig 덮어쓰기

package boot

import pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"

// resolveAuthCookieConfig applies overrides from the optional manifest
// auth.cookie sub-block. Separate helper so the Cookie-nesting stays at
// depth 2 and the parent resolver keeps a flat control shape.
func resolveAuthCookieConfig(cookie *pmanifest.CookieConfig, cfg *authInitConfig) {
	if cookie == nil {
		return
	}
	if cookie.SameSite != "" {
		cfg.SameSite = cookie.SameSite
	}
	if cookie.AccessName != "" {
		cfg.AccessName = cookie.AccessName
	}
	if cookie.RefreshName != "" {
		cfg.RefreshName = cookie.RefreshName
	}
}
