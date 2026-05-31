//ff:func feature=validate-contract type=test control=iteration dimension=1
//ff:what TestKnownCallPackages — expCalls 키의 pkg 접두사 집합 추출 검증
package contract

import (
	"testing"
)

func TestKnownCallPackages(t *testing.T) {
	exp := map[string]bool{
		"billing.Charge": true,
		"billing.Refund": true,
		"users.Find":     true,
		"NoDot":          true,
		".Bad":           true,
	}
	got := knownCallPackages(exp)
	want := map[string]bool{"billing": true, "users": true}
	if len(got) != len(want) {
		t.Fatalf("got %d packages (%v), want %d (%v)", len(got), got, len(want), want)
	}
	for k := range want {
		if !got[k] {
			t.Fatalf("missing package %q in %v", k, got)
		}
	}
	if got["NoDot"] || got[""] {
		t.Fatalf("unexpected entries in %v", got)
	}
}
