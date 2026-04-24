//ff:func feature=gen-gogin type=util control=sequence
//ff:what authInitLines — resolve 결과를 받아 main.go 에 넣을 라인 슬라이스 조립

package boot

import "fmt"

// authInitLines assembles the complete Lines slice for the auth-init block.
// Static fragments live in template_auth_init.go; dynamic pieces (TTL
// literals, secret env name, etc.) are interleaved here.
//
// Phase002 (ssac/purify) — ssac/pkg/auth is DB-free. The previous
// authInitDDLBootstrapLines (idempotent CREATE TABLE refresh_tokens) is
// removed because the DDL itself is now owned by the user's specs/db/
// directory. The `&auth.RefreshStore{DB: conn, ...}` struct literal is
// replaced by `auth.Init(infraauth.NewPostgres(queries))`, which installs
// the yongol-generated postgres RefreshStore as the package-level singleton
// consumed by auth.RefreshRotate / auth.Logout.
//
// detectReuseLogoutAll — previously a struct field on the deleted
// auth.RefreshStore; now passed as the final argument of
// auth.RefreshRotate at each call site. Nothing to wire here.
func authInitLines(cfg authInitConfig) []string {
	var out []string
	out = append(out, authInitHeaderLines...)
	out = append(out, fmt.Sprintf("accessTTL, err := time.ParseDuration(%q)", cfg.AccessTTL))
	out = append(out, authInitParseAccessTTLLines...)
	out = append(out, fmt.Sprintf("refreshTTL, err := time.ParseDuration(%q)", cfg.RefreshTTL))
	out = append(out, authInitParseRefreshTTLLines...)
	out = append(out, authInitModeOverrideLines...)
	out = append(out, fmt.Sprintf("authMode := %q", cfg.Mode))
	out = append(out, authInitModeSwitchLines...)
	out = append(out, authInitSameSiteCommentLines...)
	out = append(out, fmt.Sprintf("switch %q {", cfg.SameSite))
	out = append(out, authInitSameSiteSwitchLines...)
	out = append(out, authInitConfigureLines(cfg)...)
	out = append(out, authInitStoreInjectionLines...)
	return out
}
