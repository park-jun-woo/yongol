//ff:type feature=projectconfig type=model
//ff:what FrontendAuth — frontend.auth 섹션 모델 (token_field / refresh_field / refresh_op / store / role_field)

package manifest

// FrontendAuth mirrors the frontend.auth: section. It carries only the
// frontend-specific decisions of the auth session flow: which response
// fields hold the tokens, which operation refreshes them, and where the
// frontend stores them. It deliberately has no `mode` — the auth mode is
// a backend decision consumed via Backend.Auth.ResolvedMode() — and no
// `login_op` — which operation yields the tokens is declared by the STML
// `data-capture` attribute (plans/stml/auth-flow Phase001).
type FrontendAuth struct {
	// TokenField names the response property captured as the access
	// token (bearer mode). Required when the auth block is declared,
	// except for a role_field-only block (RoleFieldOnly) — cookie-mode
	// menu role wiring carries no token contract. XON-60 is the single
	// enforcer: it verifies the field exists in at least one OpenAPI 2xx
	// response and grants exactly the RoleFieldOnly exemption
	// (plans/stml/sitemap Phase005).
	TokenField string `yaml:"token_field"`
	// RefreshField names the response property captured as the refresh
	// token. Optional; when absent no refresh flow is generated.
	RefreshField string `yaml:"refresh_field,omitempty"`
	// RefreshOp names the operationId called on 401 to refresh the
	// access token. Optional; when absent the refresh op is inferred
	// structurally (Phase004). XON-60 verifies it when declared.
	RefreshOp string `yaml:"refresh_op,omitempty"`
	// Store selects the frontend token store: "localStorage" (default)
	// or "memory" (bearer mode only; cookie mode ignores it). SEC-404
	// rejects any other value at validate time.
	//
	// Prefer ResolvedStore() over reading Store directly — it applies
	// the default and keeps the per-call resolution in one place.
	Store string `yaml:"store,omitempty"`
	// RoleField names the auth.claims.<name> entry the sitemap
	// data-roles menu filter reads (plans/stml/sitemap Phase005).
	// Optional; required only when the sitemap uses data-roles — TM-47
	// verifies the wiring (declaration + a matching auth.claims.<name>
	// capture + non-empty backend.auth.roles).
	RoleField string `yaml:"role_field,omitempty"`
}

// ResolvedStore returns the effective frontend.auth.store after applying
// the default. An empty string in the YAML resolves to "localStorage".
// Callers that need to distinguish "explicitly set" vs "defaulted" should
// read Store directly.
func (a *FrontendAuth) ResolvedStore() string {
	if a == nil || a.Store == "" {
		return "localStorage"
	}
	return a.Store
}
