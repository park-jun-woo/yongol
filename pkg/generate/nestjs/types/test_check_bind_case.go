//ff:func feature=gen-nestjs type=test-helper control=sequence
//ff:what checkBindCase — bindCase 1 건의 Bind 결과 필드 비교 (sub-test 본체)

package types

import "testing"

// checkBindCase asserts the ir.TypeBinding produced by Bind against the
// expectations encoded in c.
func checkBindCase(t *testing.T, c bindCase) {
	t.Helper()
	reg := NewRegistry()
	b := reg.Bind(c.family, c.opts)
	if b.DBType != c.wantDBType {
		t.Errorf("DBType = %q, want %q", b.DBType, c.wantDBType)
	}
	if b.APIType != c.wantAPIType {
		t.Errorf("APIType = %q, want %q", b.APIType, c.wantAPIType)
	}
	if b.ToDBExpr != c.wantToDBExpr {
		t.Errorf("ToDBExpr = %q, want %q", b.ToDBExpr, c.wantToDBExpr)
	}
	if b.ToAPIExpr != c.wantToAPIExpr {
		t.Errorf("ToAPIExpr = %q, want %q", b.ToAPIExpr, c.wantToAPIExpr)
	}
	if b.ToResponseExpr != c.wantToRespExpr {
		t.Errorf("ToResponseExpr = %q, want %q", b.ToResponseExpr, c.wantToRespExpr)
	}
	if b.NilCheckExpr != c.wantNilCheck {
		t.Errorf("NilCheckExpr = %q, want %q", b.NilCheckExpr, c.wantNilCheck)
	}
	if b.Supported != c.wantSupported {
		t.Errorf("Supported = %v, want %v", b.Supported, c.wantSupported)
	}
	if b.Family != c.family {
		t.Errorf("Family = %v, want %v", b.Family, c.family)
	}
	if b.NotNull != c.opts.NotNull {
		t.Errorf("NotNull = %v, want %v", b.NotNull, c.opts.NotNull)
	}
}
