//ff:type feature=projectconfig type=model
//ff:what JWT 인증 설정 구조체
package manifest

type Auth struct {
	Type                   string              `yaml:"type"`       // "jwt" (required when auth is present)
	SecretEnv              string              `yaml:"secret_env"`
	RawClaims              map[string]string   `yaml:"claims"`     // YAML original: FieldName → "claim_key" or "claim_key:go_type"
	Claims                 map[string]ClaimDef `yaml:"-"`           // Parsed from RawClaims after Load()
	Roles                  []string            `yaml:"roles"`       // valid role names (e.g. ["client", "freelancer"])
	RolesLines             map[string]int      `yaml:"-"`           // role → 1-based line in manifest.yaml (0 = unknown)
	AccessTokenTTL         string              `yaml:"access_token_ttl"`         // default "15m" (Phase002)
	RefreshTokenTTL        string              `yaml:"refresh_token_ttl"`        // default "168h" (7d) (Phase002)
	DetectReuseLogoutAll   bool                `yaml:"detect_reuse_logout_all"`  // default false (Phase002)

	// Phase020 — cookie session authentication (now default).
	// Mode: "cookie" (default) | "bearer" | "hybrid". Env override:
	// BACKEND_AUTH_MODE. When Mode != "bearer" the generator emits the
	// CSRF middleware (block_csrf) and the Set-Cookie issuance path
	// (auth.SetAuthCookies / auth.ClearAuthCookies).
	//
	// Prefer ResolvedMode() over reading Mode directly — it applies the
	// Phase020 default and keeps the per-call resolution in one place.
	Mode   string        `yaml:"mode"`
	Cookie *CookieConfig `yaml:"cookie,omitempty"`
	Csrf   *CsrfConfig   `yaml:"csrf,omitempty"`
}

// ResolvedMode returns the effective auth.mode after applying the Phase020
// default. An empty string in the YAML now resolves to "cookie" (2026
// standard: HttpOnly cookie + __Host- prefix + SameSite + CSRF). Callers
// that need to distinguish "explicitly set to cookie" vs "defaulted to
// cookie" should read Mode directly.
func (a *Auth) ResolvedMode() string {
	if a == nil || a.Mode == "" {
		return "cookie"
	}
	return a.Mode
}
