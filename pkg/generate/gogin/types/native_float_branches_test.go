//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestZeroCov — 0% util/adapter 함수 (bindGoType / GoTypeOf / registry / family dispatch / pgtype) 회귀
package types

import (
	"testing"
)

func TestNativeFloat_Branches(t *testing.T) {
	nn := nativeFloat(true, "0")
	if nn.SqlcGoType != "float64" || nn.Kind != KindNative {
		t.Errorf("NOT NULL float = %+v", nn)
	}
	nullable := nativeFloat(false, "0")
	if nullable.Kind != KindPgtype {
		t.Errorf("nullable float should be pgtype, got %+v", nullable)
	}
}
