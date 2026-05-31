//ff:func feature=funcspec type=test control=sequence
//ff:what funcspec 내부 함수 by-name 직접 호출 테스트 — tsma content-aware 귀속용
package funcspec

import (
	"testing"
)

func TestFillSpecFromTypeMap_ZeroCov(t *testing.T) {
	spec := &FuncSpec{Name: "foo"}
	tm := map[string][]Field{
		"FooRequest":  {{Name: "Name"}},
		"FooResponse": {{Name: "ID"}},
	}
	fillSpecFromTypeMap(spec, tm)
	if len(spec.RequestFields) != 1 || len(spec.ResponseFields) != 1 {
		t.Errorf("spec not filled: %#v", spec)
	}
}
