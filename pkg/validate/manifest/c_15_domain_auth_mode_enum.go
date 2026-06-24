//ff:func feature=validate type=rule control=iteration dimension=1 topic=manifest-auth
//ff:what C-15 — 도메인 auth_mode 값은 SEC-403 의 validAuthModes(cookie|bearer|hybrid) 재사용 검증

package manifest

import (
	"sort"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// c15DomainAuthModeEnum rejects any domain whose auth_mode override is outside
// the closed set the generator understands. It deliberately reuses
// validAuthModes (defined in sec_403_auth_mode_enum.go) so the domain-level
// override and the backend-level backend.auth.mode share one source of truth —
// adding a new mode in one place must not silently diverge the other. Empty
// auth_mode is accepted because it means "inherit backend.auth.mode".
func c15DomainAuthModeEnum(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs == nil || fs.Manifest == nil || len(fs.Manifest.Domains) == 0 {
		return nil
	}
	names := make([]string, 0, len(fs.Manifest.Domains))
	for name := range fs.Manifest.Domains {
		names = append(names, name)
	}
	sort.Strings(names)

	var diags []diagnostic.Diagnostic
	for _, name := range names {
		mode := fs.Manifest.Domains[name].AuthMode
		if validAuthModes[mode] { // reuse SEC-403's enum (sec_403_auth_mode_enum.go)
			continue
		}
		diags = append(diags, diagnostic.Diagnostic{
			File:    "manifest.yaml",
			Phase:   diagnostic.PhaseValidate,
			Level:   diagnostic.LevelError,
			Message: "[C-15] domains." + name + ".auth_mode=\"" + mode + "\" is an unknown value",
			Advice:  "Set auth_mode to one of cookie / bearer / hybrid, or omit it to inherit backend.auth.mode.",
		})
	}
	return diags
}
