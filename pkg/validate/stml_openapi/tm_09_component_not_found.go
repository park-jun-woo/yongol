//ff:func feature=validate type=rule control=iteration dimension=1 topic=stml-openapi
//ff:what TM-09 — data-component 파일이 존재하지 않음

package stml_openapi

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// tm09Components checks that each data-component references an existing
// .tsx component file.
func tm09Components(comps []stml.ComponentRef, file string, fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, c := range comps {
		diags = append(diags, tm09Component(c.Name, file, fs)...)
	}
	return diags
}
