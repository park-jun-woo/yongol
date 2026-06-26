//ff:func feature=gen-gogin type=util control=iteration dimension=2
//ff:what collectFuncInnerTypeNames — @call 응답 내부 복합 타입 이름을 funcRespInfo 로 수집

package ssac

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/parser/funcspec"
)

// collectFuncInnerTypeNames discovers complex inner types referenced by
// @call Func Response structs. For each response type in funcRespNames it
// finds the matching FuncSpec and inspects ResponseFields for non-primitive
// types (slice-of-struct or struct references). A type is collected when:
//
//  1. It appears in needed (referenced from OpenAPI response schemas).
//  2. It is NOT a sqlc model (no DB backing row).
//  3. It is NOT already a direct @call result type in funcRespNames.
//
// Inner types inherit their parent's funcRespInfo (same package and import
// path) since they are declared in the same Func package.
func collectFuncInnerTypeNames(
	funcRespNames map[string]funcRespInfo,
	specs []funcspec.FuncSpec,
	needed map[string]bool,
	sqlcModelNames map[string]bool,
) map[string]funcRespInfo {
	result := make(map[string]funcRespInfo)
	for respName, info := range funcRespNames {
		spec := findMatchingSpec(respName, specs)
		if spec == nil {
			continue
		}
		for _, field := range spec.ResponseFields {
			bareType := strings.TrimPrefix(field.Type, "[]")
			if !needed[bareType] {
				continue
			}
			if sqlcModelNames[bareType] {
				continue
			}
			if _, exists := funcRespNames[bareType]; exists {
				continue
			}
			result[bareType] = info
		}
	}
	return result
}
