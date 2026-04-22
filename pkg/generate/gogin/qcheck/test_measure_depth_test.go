//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestMeasureDepth_Deep — 4단 중첩 fixture에서 MaxDepth≥4 검증

package qcheck

import "testing"

func TestMeasureDepth_Deep(t *testing.T) {
	reports, err := MeasureDepth("deep.go", deepSrc)
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	if len(reports) != 1 {
		t.Fatalf("want 1 report, got %d", len(reports))
	}
	if reports[0].MaxDepth < 4 {
		t.Fatalf("want depth>=4, got %d", reports[0].MaxDepth)
	}
}
