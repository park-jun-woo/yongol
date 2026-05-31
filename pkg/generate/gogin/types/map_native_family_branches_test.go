//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestZeroCov — 0% util/adapter 함수 (bindGoType / GoTypeOf / registry / family dispatch / pgtype) 회귀
package types

import (
	"testing"
)

func TestMapNativeFamily_Branches(t *testing.T) {
	if _, ok := mapNativeFamily("BIGINT", true, ""); !ok {
		t.Errorf("BIGINT should be native")
	}
	if _, ok := mapNativeFamily("REAL", true, ""); !ok {
		t.Errorf("REAL should be native")
	}
	if _, ok := mapNativeFamily("VARCHAR", true, ""); !ok {
		t.Errorf("VARCHAR should be native")
	}
	if _, ok := mapNativeFamily("BOOLEAN", true, ""); !ok {
		t.Errorf("BOOLEAN should be native")
	}
	if _, ok := mapNativeFamily("UUID", true, ""); ok {
		t.Errorf("UUID should not be native")
	}
}
