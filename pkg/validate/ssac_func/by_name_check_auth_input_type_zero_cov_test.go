//ff:func feature=validate type=test control=sequence topic=func-check
//ff:what TestByName_ZeroCov — ssac_func 헬퍼들을 이름으로 직접 호출해 커버리지 귀속
package ssac_func

import (
	"testing"

	parsessac "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestByNameCheckAuthInputType_ZeroCov(t *testing.T) {
	// Empty Inputs → no per-input resolution; function is exercised by name.
	fn := parsessac.ServiceFunc{Name: "GetItem", FileName: "f.ssac"}
	seq := parsessac.Sequence{Line: 5, Inputs: map[string]string{}}
	if d := checkAuthInputType(nil, fn, seq); len(d) != 0 {
		t.Errorf("checkAuthInputType empty inputs should yield no diagnostics, got %v", d)
	}
}
