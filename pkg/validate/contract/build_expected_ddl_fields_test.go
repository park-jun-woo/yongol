//ff:func feature=validate-contract type=test control=iteration dimension=1
//ff:what TestBuildExpectedDDLFields — DDL 컬럼명을 canonical(소문자·무언더스코어) 키로 집합화 검증

package contract

import "testing"

func TestBuildExpectedDDLFields(t *testing.T) {
	fs := buildFSForPRV02() // users: id / email / created_at
	got := buildExpectedDDLFields(fs)
	// created_at canonicalizes to "createdat"; id → "id"; email → "email".
	want := []string{canonicalFieldKey("id"), canonicalFieldKey("email"), canonicalFieldKey("created_at")}
	for _, k := range want {
		if !got[k] {
			t.Errorf("expected canonical key %q present in %v", k, got)
		}
	}
	if got[canonicalFieldKey("created_at")] != got["createdat"] {
		t.Error("created_at should normalize to createdat")
	}
}
