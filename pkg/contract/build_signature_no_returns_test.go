//ff:func feature=contract type=test control=sequence
//ff:what test: TestBuildSignature — FuncDecl→FuncSignature 변환, error 반환 시 HasErr, 반환 없음 분기 검증
package contract

import (
	"testing"
)

func TestBuildSignatureNoReturns(t *testing.T) {
	sig := buildSigFromSrc(t, "func Run() {}")
	if sig.Name != "Run" {
		t.Errorf("name: got %q want Run", sig.Name)
	}
	if len(sig.Returns) != 0 {
		t.Errorf("returns: got %v want none", sig.Returns)
	}
	if sig.HasErr {
		t.Errorf("expected HasErr false")
	}
}
