//ff:func feature=gen-gogin type=test-helper control=sequence
//ff:what checkMatrixCase — matrixCase 1 건의 binding 필드 비교 (sub-test 본체)

package types

import (
	"strings"
	"testing"
)

// checkMatrixCase asserts the binding produced by MapPGType against the
// expectations encoded in c. Used as the body of every sub-test so the
// per-family Test functions stay within the F1 / Q4 line budget.
func checkMatrixCase(t *testing.T, c matrixCase) {
	t.Helper()
	b := MapPGType(c.col)
	if b.SqlcGoType != c.wantSqlcType {
		t.Errorf("SqlcGoType = %q, want %q", b.SqlcGoType, c.wantSqlcType)
	}
	if b.ApiField != c.wantApiField {
		t.Errorf("ApiField = %q, want %q", b.ApiField, c.wantApiField)
	}
	if b.Kind != c.wantKind {
		t.Errorf("Kind = %v, want %v", b.Kind, c.wantKind)
	}
	if b.NeedsOverride != c.wantNeedsOver {
		t.Errorf("NeedsOverride = %v, want %v", b.NeedsOverride, c.wantNeedsOver)
	}
	if b.Supported != c.wantSupported {
		t.Errorf("Supported = %v, want %v", b.Supported, c.wantSupported)
	}
	if c.wantConvertSub != "" && !strings.Contains(b.ConvertExpr, c.wantConvertSub) {
		t.Errorf("ConvertExpr = %q, want substring %q", b.ConvertExpr, c.wantConvertSub)
	}
}
