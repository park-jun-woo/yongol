//ff:func feature=validate-contract type=util control=sequence
//ff:what compareSignature — 두 FuncSignature 비교해 차이 설명 문자열 목록 반환

package contract

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/contract"
)

// compareSignature compares an expected (SSOT-derived) signature with a
// preserved one and returns a list of human-readable difference
// descriptions. An empty slice means the signatures match exactly.
//
// The comparison is intentionally strict: any rename, type change, or
// arity difference surfaces. The only tolerated variance is parameter
// / return naming (Go does not require return names to match on the
// caller side).
func compareSignature(expected, actual contract.FuncSignature) []string {
	var diffs []string
	if expected.Name != actual.Name {
		diffs = append(diffs, fmt.Sprintf("function name: expected %q, got %q", expected.Name, actual.Name))
	}
	if len(expected.Params) != len(actual.Params) {
		diffs = append(diffs, fmt.Sprintf("parameter arity: expected %d, got %d", len(expected.Params), len(actual.Params)))
	} else {
		diffs = append(diffs, compareParams(expected.Params, actual.Params)...)
	}
	if len(expected.Returns) != len(actual.Returns) {
		diffs = append(diffs, fmt.Sprintf("return arity: expected %d, got %d", len(expected.Returns), len(actual.Returns)))
	} else {
		diffs = append(diffs, compareReturns(expected.Returns, actual.Returns)...)
	}
	if expected.HasErr != actual.HasErr {
		diffs = append(diffs, fmt.Sprintf("error-return: expected %v, got %v", expected.HasErr, actual.HasErr))
	}
	return diffs
}
