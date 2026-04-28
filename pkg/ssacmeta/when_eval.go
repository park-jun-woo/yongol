//ff:func feature=ssacmeta type=util control=sequence
//ff:what EvaluateWhen — interface.yaml 의 when: 표현식을 manifest 컨텍스트로 평가

package ssacmeta

import "strings"

// EvaluateWhen returns true when the port should be active under the given
// manifest context. Supported syntax (minimal):
//
//	always                              — always active
//	manifest.<path> == "<value>"        — equal comparison on a manifest field
//	manifest.<path>                     — truthy check on a manifest field
//
// Paths are dot-separated keys into the flattened manifest map. Unknown
// paths evaluate to false.
func EvaluateWhen(expr string, manifest map[string]any) bool {
	e := strings.TrimSpace(expr)
	if e == "" || e == "always" {
		return true
	}
	if !strings.HasPrefix(e, "manifest.") {
		return false
	}
	if i := strings.Index(e, "=="); i > 0 {
		return evaluateWhenEquality(e, i, manifest)
	}
	return evaluateWhenTruthy(e, manifest)
}
