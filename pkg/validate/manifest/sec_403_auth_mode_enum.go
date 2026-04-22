//ff:func feature=validate type=rule control=sequence topic=manifest-auth
//ff:what SEC-403 — backend.auth.mode must be one of cookie|bearer|hybrid (Phase020)

package manifest

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// validAuthModes is the closed set of auth.mode values the generator
// understands. Empty string is also accepted because it resolves to
// "cookie" via Auth.ResolvedMode() — authors that omit the key entirely
// should not be forced into spelling out the default.
var validAuthModes = map[string]bool{
	"":       true, // resolves to "cookie" default
	"cookie": true,
	"bearer": true,
	"hybrid": true,
}

// sec403AuthModeEnum rejects manifests whose backend.auth.mode value is
// outside {cookie, bearer, hybrid}. A typo like "cookies" or "jwt" today
// silently falls through to bearer-like behaviour; Phase020 elevates
// unknown values to a hard ERROR so the operator sees the problem at
// validate time instead of runtime.
func sec403AuthModeEnum(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs == nil || fs.Manifest == nil || fs.Manifest.Backend.Auth == nil {
		return nil
	}
	mode := fs.Manifest.Backend.Auth.Mode
	if validAuthModes[mode] {
		return nil
	}
	return []diagnostic.Diagnostic{{
		File:    "manifest.yaml",
		Phase:   diagnostic.PhaseValidate,
		Level:   diagnostic.LevelError,
		Message: "[SEC-403] backend.auth.mode=\"" + mode + "\" is an unknown value",
		Advice:  "Set auth.mode to one of cookie / bearer / hybrid (defaults to cookie when omitted)",
	}}
}
