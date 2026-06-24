//ff:func feature=generate type=util control=iteration dimension=1
//ff:what DomainAuthFor — 도메인별 Auth 파생 (auth_mode override + backend 상속, 멀티 도메인 전용)

package prepared

import "github.com/park-jun-woo/yongol/pkg/yongol"

// DomainAuthFor resolves one prepared.Auth per manifest domain. It starts from
// the backend-wide AuthFor(fs) derivation and, when a domain declares its own
// auth_mode, overrides Mode (and the derived CsrfRequired flag) for that
// domain. Domains without an auth_mode inherit the backend mode unchanged.
//
// Returns nil for single-site projects (no domains). The backend claims/TTL/
// cookie/csrf config in Raw stays shared across domains — only the transport
// mode diverges per domain (Phase008 §3c).
func DomainAuthFor(fs *yongol.Fullstack) map[string]Auth {
	if fs == nil || fs.Manifest == nil || len(fs.Manifest.Domains) == 0 {
		return nil
	}
	base := AuthFor(fs)
	out := make(map[string]Auth, len(fs.Manifest.Domains))
	for name, cfg := range fs.Manifest.Domains {
		a := base
		if cfg.AuthMode != "" {
			a.Mode = cfg.AuthMode
			a.CsrfRequired = a.Mode == "cookie" || a.Mode == "hybrid"
		}
		out[name] = a
	}
	return out
}
