//ff:func feature=validate-contract type=test control=sequence
//ff:what TestKnownCallPackages — expCalls 키의 pkg 접두사 집합 추출 검증
package contract

import (
	"testing"
)

func TestKnownCallPackagesEmpty(t *testing.T) {
	if got := knownCallPackages(map[string]bool{}); len(got) != 0 {
		t.Fatalf("expected empty, got %v", got)
	}
}
