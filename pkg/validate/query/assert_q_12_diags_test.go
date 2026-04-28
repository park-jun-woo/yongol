//ff:func feature=validate type=test-helper control=iteration dimension=1 topic=query-structural
//ff:what assertQ12Diags — Q-12 진단 개수 / 메시지 / advice 내용 검증

package query

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

// assertQ12Diags checks that the diagnostics emitted for one Q-12 case
// match the expected count, that each required substring appears in the
// first message, and that the advice carries the pgx/v5 pgtype import
// (whenever a diagnostic was expected).
func assertQ12Diags(t *testing.T, tc q12UuidTestCase, diags []diagnostic.Diagnostic) {
	t.Helper()
	if len(diags) != tc.wantDiags {
		t.Fatalf("diag count: want %d, got %d (%+v)", tc.wantDiags, len(diags), diags)
	}
	if tc.wantDiags == 0 {
		return
	}
	for _, sub := range tc.wantMsgSubstrs {
		if !strings.Contains(diags[0].Message, sub) {
			t.Errorf("message missing %q: %q", sub, diags[0].Message)
		}
	}
	if !strings.Contains(diags[0].Advice, "github.com/jackc/pgx/v5/pgtype") {
		t.Errorf("advice missing pgtype import: %q", diags[0].Advice)
	}
}
