//ff:func feature=validate type=util control=sequence topic=stml-openapi
//ff:what xmo11Diag — XMO-11 (프론트엔드 ON 인데 STML 0개) 진단 리터럴 생성 (단일/도메인 공유)

package stml_openapi

import "github.com/park-jun-woo/yongol/pkg/diagnostic"

// xmo11Diag returns the single XMO-11 ERROR. Shared by the single-site
// xmo11NoStml and the domain-mode xmo11NoStmlAll so the message stays in one
// place.
func xmo11Diag() []diagnostic.Diagnostic {
	return []diagnostic.Diagnostic{{
		File:    "manifest.yaml",
		Phase:   diagnostic.PhaseValidate,
		Level:   diagnostic.LevelError,
		Message: "[XMO-11] frontend is enabled but no STML pages were found",
		Advice:  "Author STML pages under specs/frontend/, or set frontend.enabled: false in manifest.yaml if this project has no frontend",
	}}
}
