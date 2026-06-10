//ff:func feature=stml-gen type=util control=sequence
//ff:what ParamBind의 Source를 JSX 표현식으로 변환한다
package stml

import (
	"strings"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func paramSourceExpr(p stmlparser.ParamBind) string {
	if strings.HasPrefix(p.Source, "route.") {
		return strings.TrimPrefix(p.Source, "route.")
	}
	if strings.HasPrefix(p.Source, "item.") {
		// Row-context source (page-flow Phase006): valid as-is inside the
		// data-each map callback scope ((item) => ...).
		return p.Source
	}
	return p.Source
}
