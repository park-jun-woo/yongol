//ff:func feature=stml-gen type=util control=iteration dimension=1 topic=import-collect
//ff:what FetchBlock에서 필요한 임포트(useParams, 컴포넌트)를 수집한다
package stml

import (
	"strings"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func collectFetchImports(f stmlparser.FetchBlock, is *importSet, compSet map[string]bool) {
	for _, p := range f.Params {
		if strings.HasPrefix(p.Source, "route.") {
			is.useParams = true
		}
	}
	for _, c := range f.Components {
		compSet[c.Name] = true
	}
	if len(f.Eaches) > 0 {
		is.useTable = true
	}
	for _, child := range f.NestedFetches {
		collectFetchImports(child, is, compSet)
	}
}
