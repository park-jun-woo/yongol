//ff:func feature=manifest type=accessor control=selection
//ff:what ResolveRef — SSaC manifest.* 참조 경로를 실제 값으로 해석 (화이트리스트 방식)

package manifest

// ResolveRef resolves a manifest reference path (the part after "manifest.")
// against the given ProjectConfig. It returns the resolved value and true if
// the path is known and the value is present, or zero value and false otherwise.
//
// Supported paths (whitelist):
//   - auth.accessTokenTTL  → backend.auth.access_token_ttl (duration → seconds int64)
//   - auth.refreshTokenTTL → backend.auth.refresh_token_ttl (duration → seconds int64)
func ResolveRef(cfg *ProjectConfig, path string) (RefValue, bool) {
	if cfg == nil {
		return RefValue{}, false
	}
	switch path {
	case "auth.accessTokenTTL":
		return resolveDurationTTL(cfg.Backend.Auth, func(a *Auth) string { return a.AccessTokenTTL })
	case "auth.refreshTokenTTL":
		return resolveDurationTTL(cfg.Backend.Auth, func(a *Auth) string { return a.RefreshTokenTTL })
	default:
		return RefValue{}, false
	}
}
