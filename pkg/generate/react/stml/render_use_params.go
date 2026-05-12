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

// renderUseParamsWithRoute generates useParams destructuring by merging
// param names from data-param-* bindings and from the data-route pattern.
// Route pattern params (e.g. ":buildingId" from "/buildings/:buildingId/units/:id")
// are included even if no ParamBind references them.
func renderUseParamsWithRoute(params []stmlparser.ParamBind, route string) string {
	routeParams := mergeRouteParamNames(
		extractRouteParamNames(params),
		extractRoutePatternParams(route),
	)
	if len(routeParams) == 0 {
		return ""
	}
	return fmt.Sprintf("const { %s } = useParams()", strings.Join(routeParams, ", "))
}
