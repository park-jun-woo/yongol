//ff:func feature=stml-gen type=util control=selection dimension=1
//ff:what GuardLifecycle 노드(ref.loading|error|empty)를 JSX 표현식으로 변환한다
package stml

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// guardLifecycleToJSX renders a fetch lifecycle node, mirroring the legacy
// switch output: loading → "<dataVar>Loading", error → "<dataVar>Error",
// empty → "<dataVar>.<path>?.length === 0".
func guardLifecycleToJSX(expr *stml.GuardExpr, dataVar string) string {
	switch expr.Lifecycle {
	case "loading":
		return dataVar + "Loading"
	case "error":
		return dataVar + "Error"
	default:
		return fmt.Sprintf("%s.%s?.length === 0", dataVar, optionalChainPath(expr.Ref.Path()))
	}
}
