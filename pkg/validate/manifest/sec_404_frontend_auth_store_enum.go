//ff:func feature=validate type=rule control=sequence topic=manifest-auth
//ff:what SEC-404 — frontend.auth.store must be one of localStorage|memory (stml/auth-flow Phase001)

package manifest

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// validFrontendAuthStores is the closed set of frontend.auth.store values
// the generator understands. Empty string is also accepted because it
// resolves to "localStorage" via FrontendAuth.ResolvedStore() — authors
// that omit the key should not be forced into spelling out the default.
var validFrontendAuthStores = map[string]bool{
	"":             true, // resolves to "localStorage" default
	"localStorage": true,
	"memory":       true,
}

// sec404FrontendAuthStoreEnum rejects manifests whose frontend.auth.store
// value is outside {localStorage, memory}. "cookie" is a frequent mistake:
// cookies are an auth *mode* (backend.auth.mode), not a frontend token
// store — in cookie mode the browser holds httpOnly cookies and there is
// nothing for the frontend to store, so the value is rejected with a
// pointer to backend.auth.mode instead.
func sec404FrontendAuthStoreEnum(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs == nil || fs.Manifest == nil || fs.Manifest.Frontend.Auth == nil {
		return nil
	}
	store := fs.Manifest.Frontend.Auth.Store
	if validFrontendAuthStores[store] {
		return nil
	}
	advice := "Set frontend.auth.store to one of localStorage / memory (defaults to localStorage when omitted)"
	if store == "cookie" || store == "cookies" {
		advice = "cookie is an auth mode, not a frontend store — set `backend.auth.mode: cookie` and remove frontend.auth.store (httpOnly cookies leave nothing for the frontend to store)"
	}
	return []diagnostic.Diagnostic{{
		File:    "manifest.yaml",
		Phase:   diagnostic.PhaseValidate,
		Level:   diagnostic.LevelError,
		Message: "[SEC-404] frontend.auth.store=\"" + store + "\" is an unknown value",
		Advice:  advice,
	}}
}
