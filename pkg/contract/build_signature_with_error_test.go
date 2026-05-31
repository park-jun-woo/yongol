//ff:func feature=contract type=test control=sequence
//ff:what test: TestBuildSignature — FuncDecl→FuncSignature 변환, error 반환 시 HasErr, 반환 없음 분기 검증
package contract

import (
	"testing"
)

func TestBuildSignatureWithError(t *testing.T) {
	sig := buildSigFromSrc(t, "func Do(id int64) (string, error) {}")
	if sig.Name != "Do" {
		t.Errorf("name: got %q want Do", sig.Name)
	}
	if len(sig.Params) != 1 || sig.Params[0] != (FuncParam{Name: "id", Type: "int64"}) {
		t.Errorf("params: got %+v", sig.Params)
	}
	if len(sig.Returns) != 2 || sig.Returns[0] != "string" || sig.Returns[1] != "error" {
		t.Errorf("returns: got %v", sig.Returns)
	}
	if !sig.HasErr {
		t.Errorf("expected HasErr true")
	}
}
