//ff:func feature=stml-gen type=generator control=sequence
//ff:what route 파라미터의 useParams 구조분해 할당 코드를 생성한다
package stml

import (
	"fmt"
	"strings"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// renderUseParams generates useParams destructuring for route params.
func renderUseParams(params []stmlparser.ParamBind) string {
	routeParams := extractRouteParamNames(params)
	if len(routeParams) == 0 {
		return ""
	}
	return fmt.Sprintf("const { %s } = useParams()", strings.Join(routeParams, ", "))
}
