//ff:func feature=validate type=rule control=sequence topic=stml-openapi
//ff:what XMO-11 — Frontend ON인데 STML 페이지가 0개 (ERROR, gozhip 구멍 차단)

package stml_openapi

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xmo11NoStml emits a single ERROR when the frontend is ON but no STML pages
// exist. A declared frontend with zero pages is an unfinished or
// misconfigured project rather than a backend-only one; XMO-11 closes the gap
// where 0-page projects previously generated without any coverage diagnostic.
func xmo11NoStml(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if !frontendEnabled(fs) || len(fs.STMLPages) != 0 {
		return nil
	}
	return xmo11Diag()
}
