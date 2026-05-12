//ff:func feature=validate type=util control=sequence topic=stml-design
//ff:what recordComponent — data-component 참조 기록
package stml_design

import (
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// recordComponent records a data-component reference.
func recordComponent(comp stml.ComponentRef, file string, out *pageTokenRefs) {
	if comp.Name != "" {
		out.Components = append(out.Components, tokenRef{
			File: file,
			Name: comp.Name,
		})
	}
}
