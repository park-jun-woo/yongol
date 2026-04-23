//ff:func feature=gen-gogin type=util control=sequence
//ff:what authInitLines — resolve 결과를 받아 main.go 에 넣을 라인 슬라이스 조립

package boot

import "fmt"

// authInitLines assembles the complete Lines slice for the auth-init block.
// Static fragments live in template_auth_init.go; dynamic pieces (TTL
// literals, secret env name, etc.) are interleaved here.
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
	out = append(out, authInitDDLBootstrapLines...)
	out = append(out, fmt.Sprintf("refreshStore := &auth.RefreshStore{DB: conn, DetectReuseLogoutAll: %t}", cfg.DetectReuse))
	out = append(out, authInitStoreInjectionLines...)
	return out
}
