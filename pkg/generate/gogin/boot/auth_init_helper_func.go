//ff:func feature=gen-gogin type=util control=sequence
//ff:what authInitHelperFunc — cmd/configure_auth.go 에 emit 할 configureAuth 함수 본문 생성

package boot

import "fmt"

// authInitHelperFunc returns the configureAuth top-level func declaration
// emitted into cmd/configure_auth.go. The func encapsulates TTL parsing,
// auth mode resolution, SameSite selection, and auth.Configure — all of
// which were previously inlined in main() and pushed it past Q3.
func authInitHelperFunc(cfg authInitConfig) string {
	return fmt.Sprintf(`func configureAuth(accessTTLStr, refreshTTLStr, defaultMode, sameSiteStr, secretEnv string) {
	accessTTL, err := time.ParseDuration(accessTTLStr)
	if err != nil {
		slog.Error("parse access_token_ttl", "err", err)
		os.Exit(1)
	}
	refreshTTL, err := time.ParseDuration(refreshTTLStr)
	if err != nil {
		slog.Error("parse refresh_token_ttl", "err", err)
		os.Exit(1)
	}
	authMode := defaultMode
	if v := os.Getenv("BACKEND_AUTH_MODE"); v != "" {
		switch v {
		case "bearer", "cookie", "hybrid":
			authMode = v
		}
	}
	sameSite := parseSameSite(sameSiteStr)
	auth.Configure(auth.Config{
		SecretEnv:  secretEnv,
		AccessTTL:  accessTTL,
		RefreshTTL: refreshTTL,
		Mode:       authMode,
		CookieAttrs: auth.CookieAttrs{
			AccessName:  %q,
			RefreshName: %q,
			SameSite:    sameSite,
			AccessTTL:   accessTTL,
			RefreshTTL:  refreshTTL,
		},
	})
}`, cfg.AccessName, cfg.RefreshName)
}
