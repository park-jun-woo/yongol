//ff:func feature=validate-contract type=test control=sequence
//ff:what TestBuildExpectedCalls — FuncSpec + SSaC.callRef 에서 허용 호출 대상 집합 구축 검증

package contract

import "testing"

func TestBuildExpectedCalls(t *testing.T) {
	fs := buildFSForPRV02() // Ground SSaC.callRef = billing.checkCredits
	g := fs.Ground()
	calls := buildExpectedCalls(fs, g)
	if !calls["billing.checkCredits"] {
		t.Error("expected verbatim SSaC call ref")
	}
	if !calls["billing.CheckCredits"] {
		t.Error("expected PascalCase variant of SSaC call ref")
	}

	t.Run("nil ground tolerated", func(t *testing.T) {
		c := buildExpectedCalls(fs, nil)
		if c == nil {
			t.Fatal("expected non-nil map even with nil ground")
		}
	})
}
