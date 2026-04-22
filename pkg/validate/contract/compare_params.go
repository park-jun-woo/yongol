//ff:func feature=validate-contract type=util control=iteration dimension=1
//ff:what compareParams — 파라미터 타입 불일치를 인덱스 단위 문자열 목록으로 반환

package contract

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/contract"
)

// compareParams assumes both slices share equal length (caller checks
// arity first) and returns one diff per mismatching type. Parameter
// names are ignored because Go permits callers to omit names in their
// call sites — only types carry the call contract.
func compareParams(expected, actual []contract.FuncParam) []string {
	var diffs []string
	for i, p := range expected {
		if p.Type != actual[i].Type {
			diffs = append(diffs, fmt.Sprintf("param #%d type: expected %q, got %q", i+1, p.Type, actual[i].Type))
		}
	}
	return diffs
}
