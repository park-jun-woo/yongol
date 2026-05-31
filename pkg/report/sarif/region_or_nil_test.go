//ff:func feature=report type=test control=sequence topic=sarif
//ff:what TestRegionOrNil — 양수 line 은 *Region, 0/음수 line 은 nil 반환 검증
package sarif

import (
	"testing"
)

func TestRegionOrNil(t *testing.T) {
	if r := regionOrNil(0); r != nil {
		t.Errorf("line 0: got %+v, want nil", r)
	}
	if r := regionOrNil(-3); r != nil {
		t.Errorf("line -3: got %+v, want nil", r)
	}
	r := regionOrNil(42)
	if r == nil {
		t.Fatalf("line 42: got nil, want Region")
	}
	if r.StartLine != 42 {
		t.Errorf("StartLine: got %d, want 42", r.StartLine)
	}
	if r.StartColumn != 0 {
		t.Errorf("StartColumn: got %d, want 0", r.StartColumn)
	}
}
