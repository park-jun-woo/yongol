//ff:func feature=validate-contract type=util control=iteration dimension=1
//ff:what compareReturns — 반환 타입 불일치를 인덱스 단위 문자열 목록으로 반환

package contract

import "fmt"

// compareReturns assumes both slices share equal length and returns one
// diff per mismatching return type. The comparison is purely textual
// since ExtractSignature renders return types via go/printer — that
// stable rendering is enough to spot user-visible contract changes.
func compareReturns(expected, actual []string) []string {
	var diffs []string
	for i, r := range expected {
		if r != actual[i] {
			diffs = append(diffs, fmt.Sprintf("return #%d type: expected %q, got %q", i+1, r, actual[i]))
		}
	}
	return diffs
}
