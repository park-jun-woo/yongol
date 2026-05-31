//ff:func feature=gen-typemap type=test control=sequence
//ff:what typemap 내부 함수 by-name 직접 호출 테스트 — tsma content-aware 귀속용
package typemap

import (
	"testing"
)

func TestClassifyPgtypeFamily_ZeroCov(t *testing.T) {
	if fam, ok := classifyPgtypeFamily("UUID"); !ok || fam != FamilyUUID {
		t.Errorf("UUID -> %v ok=%v", fam, ok)
	}
	if fam, ok := classifyPgtypeFamily("JSONB"); !ok || fam != FamilyJSONB {
		t.Errorf("JSONB -> %v ok=%v", fam, ok)
	}
	if _, ok := classifyPgtypeFamily("NOPE"); ok {
		t.Error("unknown should be false")
	}
}
