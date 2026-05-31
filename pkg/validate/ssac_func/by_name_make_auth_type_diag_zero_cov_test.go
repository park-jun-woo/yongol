//ff:func feature=validate type=test control=sequence topic=func-check
//ff:what TestByName_ZeroCov — ssac_func 헬퍼들을 이름으로 직접 호출해 커버리지 귀속
package ssac_func

import (
	"testing"
)

func TestByNameMakeAuthTypeDiag_ZeroCov(t *testing.T) {
	// string-compatible source → nil.
	if d := makeAuthTypeDiag("f.ssac", 10, "id", "string", "Op"); d != nil {
		t.Errorf("makeAuthTypeDiag string-compatible should be nil")
	}
	// empty source type → nil.
	if d := makeAuthTypeDiag("f.ssac", 10, "id", "", "Op"); d != nil {
		t.Errorf("makeAuthTypeDiag empty should be nil")
	}
	// incompatible (uuid) → diagnostic.
	if d := makeAuthTypeDiag("f.ssac", 10, "id", "uuid.UUID", "Op"); d == nil {
		t.Errorf("makeAuthTypeDiag incompatible should produce diagnostic")
	}
}
