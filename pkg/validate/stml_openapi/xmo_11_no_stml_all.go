//ff:func feature=validate type=rule control=sequence topic=stml-openapi
//ff:what XMO-11 (도메인) — Frontend ON인데 전체 도메인 통틀어 STML 페이지가 0개 (ERROR)

package stml_openapi

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xmo11NoStmlAll is the domain-mode XMO-11: it fires the same ERROR when the
// frontend is ON but no STML pages exist in ANY domain. fs.AllSTMLPages()
// flattens every domain's pages, so a project that declares domains yet ships
// zero pages anywhere is still caught.
func xmo11NoStmlAll(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if !frontendEnabled(fs) || len(fs.AllSTMLPages()) != 0 {
		return nil
	}
	return xmo11Diag()
}
