//ff:func feature=rule type=loader control=iteration dimension=1
//ff:what populatePathOpsParams — path의 각 operation에서 param/sort/filter 등록
package ground

import (
	"github.com/getkin/kin-openapi/openapi3"

	"github.com/park-jun-woo/yongol/pkg/rule"
)

func populatePathOpsParams(g *rule.Ground, ops map[string]*openapi3.Operation) {
	for _, op := range ops {
		// operationId 부재 operation 은 Lookup 키를 만들 수 없어 skip.
		// operationId 는 OpenAPI 사양상 선택이지만 yongol 는 전 SSOT 연결
		// 키로 사용하므로 선행 validate (`pkg/validate/openapi` 의 O-4 규칙)
		// 에서 ERROR 로 차단된다. Ground 계층은 loader 이므로 diagnostic 을
		// 발행하지 않고 skip 만 한다 (O-4 통과 후에는 실제로 도달하지 않음).
		if op.OperationID != "" {
			populateOpParams(g, op)
		}
	}
}
