//ff:func feature=gen-gogin type=util control=sequence topic=csrf
//ff:what csrfActive — auth.mode=cookie|hybrid 이고 csrf.enabled=true 인지 판정

package middleware

import "github.com/park-jun-woo/yongol/pkg/yongol"

// csrfActive mirrors the boot.blockCsrf Active condition so the
// middleware file is emitted on the same trigger. Kept local to avoid
// import cycles (middleware → boot).
func csrfActive(fs *yongol.Fullstack) bool {
	if fs == nil || fs.Manifest == nil || fs.Manifest.Backend.Auth == nil {
		return false
	}
	a := fs.Manifest.Backend.Auth
	mode := a.Mode
	if mode != "cookie" && mode != "hybrid" {
		return false
	}
	if a.Csrf == nil {
		// Default on for cookie/hybrid modes — SEC-201 rejects the
		// explicit-false combination at validate time; reaching codegen
		// with nil Csrf means "accept defaults, enabled".
		return true
	}
	return a.Csrf.Enabled
}
