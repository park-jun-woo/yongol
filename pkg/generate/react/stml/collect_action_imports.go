//ff:func feature=stml-gen type=util control=iteration dimension=1 topic=import-collect
//ff:what ActionBlock에서 필요한 임포트(useParams, 컴포넌트)를 수집한다
package stml

import (
	"strings"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func collectActionImports(a stmlparser.ActionBlock, is *importSet, compSet map[string]bool) {
	for _, p := range a.Params {
		if strings.HasPrefix(p.Source, "route.") {
			is.useParams = true
		}
	}
	for _, f := range a.Fields {
		if strings.HasPrefix(f.Tag, "data-component:") {
			comp := strings.TrimPrefix(f.Tag, "data-component:")
			compSet[comp] = true
		}
	}
}
