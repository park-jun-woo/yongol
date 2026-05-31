//ff:func feature=contract type=test control=sequence
//ff:what test: TestBuildSignature — FuncDecl→FuncSignature 변환, error 반환 시 HasErr, 반환 없음 분기 검증
package contract

import (
	"testing"
)

func TestBuildSignatureNonErrorReturn(t *testing.T) {
	sig := buildSigFromSrc(t, "func Count() int { return 0 }")
	if len(sig.Returns) != 1 || sig.Returns[0] != "int" {
		t.Errorf("returns: got %v want [int]", sig.Returns)
	}
	if sig.HasErr {
		t.Errorf("expected HasErr false for int return")
	}
}
