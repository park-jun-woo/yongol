//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestNilCheckExpr — NotNull 분기 커버
package types

import "testing"

func TestNilCheckExpr_ZeroCov(t *testing.T) {
	if nilCheckExpr(true) != "" {
		t.Errorf("NOT NULL should give empty nil check")
	}
	if nilCheckExpr(false) != "{var} === null" {
		t.Errorf("nullable nil check wrong")
	}
}
