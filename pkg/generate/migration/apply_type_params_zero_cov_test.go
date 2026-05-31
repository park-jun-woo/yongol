//ff:func feature=migration type=test control=sequence
//ff:what 마이그레이션 내부 함수 by-name 직접 호출 테스트 — tsma content-aware 귀속용
package migration

import (
	"testing"
)

func TestApplyTypeParams_ZeroCov(t *testing.T) {
	ct := &CanonicalType{Base: "NUMERIC"}
	applyTypeParams(ct, "10,2")
	if ct.Precision != 10 || ct.Scale != 2 {
		t.Errorf("numeric params wrong: %#v", ct)
	}
	ct2 := &CanonicalType{Base: "VARCHAR"}
	applyTypeParams(ct2, "255")
	if ct2.Length != 255 {
		t.Errorf("varchar length wrong")
	}
}
