//ff:func feature=validate type=test control=sequence dimension=1 topic=stml-openapi
//ff:what TM-17 단일 케이스 진단 기대치 검증 헬퍼

package stml_openapi

import (
	"strings"
	"testing"
)

// assertTM17 runs tm17GuardSyntax on one condition and asserts whether a
// diagnostic (with the [TM-17] prefix) is expected.
func assertTM17(t *testing.T, condition string, wantDiag bool) {
	t.Helper()
	diags := tm17GuardSyntax(condition, "page.html")
	if wantDiag && len(diags) == 0 {
		t.Fatalf("expected a TM-17 diagnostic for %q, got none", condition)
	}
	if !wantDiag && len(diags) != 0 {
		t.Fatalf("expected no diagnostic for %q, got %+v", condition, diags)
	}
	if wantDiag && !strings.Contains(diags[0].Message, "[TM-17]") {
		t.Fatalf("expected [TM-17] prefix, got %q", diags[0].Message)
	}
}
