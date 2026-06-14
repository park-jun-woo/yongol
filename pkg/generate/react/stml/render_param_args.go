//ff:func feature=stml-gen type=util control=iteration dimension=1
//ff:what ParamBind 슬라이스에서 { key: value } 형태의 인자 문자열을 생성한다
package stml

import (
	"fmt"
	"strings"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func renderParamArgs(params []stmlparser.ParamBind, opID string, pathParamTypes map[string]map[string]string) string {
	if len(params) == 0 {
		return ""
	}
	var parts []string
	for _, p := range params {
		expr := paramSourceExpr(p)
		if isIntegerParam(opID, p.Name, pathParamTypes) {
			expr = wrapNumberArg(expr)
		}
		parts = append(parts, fmt.Sprintf("%s: %s", p.Name, expr))
	}
	return "{ " + strings.Join(parts, ", ") + " }"
}
