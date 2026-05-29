//ff:func feature=orchestrator type=test control=sequence
//ff:what DetectSSOTs — api/ 디렉토리 부재 시 KindOpenAPI 미포함
package yongol

import (
	"testing"
)

// TestDetectSSOTsOpenAPIAbsent confirms that a bare specs root with no api/
// directory produces no KindOpenAPI entry at all (SSOTAbsent is omitted).
func TestDetectSSOTsOpenAPIAbsent(t *testing.T) {
	tmp := newTmpSpecsDir(t)

	detected, err := DetectSSOTs(tmp)
	if err != nil {
		t.Fatalf("DetectSSOTs: %v", err)
	}
	if _, ok := hasKind(detected, KindOpenAPI); ok {
		t.Fatalf("KindOpenAPI unexpectedly detected; detected=%+v", detected)
	}
}
