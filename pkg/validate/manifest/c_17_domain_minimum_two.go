//ff:func feature=validate type=rule control=sequence topic=manifest-structural
//ff:what C-17 — domains 선언 시 최소 2개 도메인 필수 (1개면 단일 사이트 사용) ERROR

package manifest

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// c17DomainMinimumTwo rejects a `domains:` block that declares exactly one
// domain. The multi-site machinery (per-domain route groups, cross-domain
// security rules, separate frontends) only earns its complexity with two or
// more domains; a single domain should be expressed as a plain single-site
// project (top-level openapi + frontend) instead. An absent or empty domains
// block is the single-site path and is left untouched.
func c17DomainMinimumTwo(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs == nil || fs.Manifest == nil || len(fs.Manifest.Domains) == 0 {
		return nil
	}
	if len(fs.Manifest.Domains) >= 2 {
		return nil
	}
	return []diagnostic.Diagnostic{{
		File:    "manifest.yaml",
		Phase:   diagnostic.PhaseValidate,
		Level:   diagnostic.LevelError,
		Message: "[C-17] domains declares only 1 domain; multi-site requires at least 2",
		Advice:  "Add a second domain, or remove the domains block and use a single-site project (top-level openapi + frontend).",
	}}
}
