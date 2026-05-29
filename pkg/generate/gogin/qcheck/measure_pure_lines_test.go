//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestMeasurePureLines_Long — range body 11 순수 라인 fixture에서 PureLines>10 검증

package qcheck

import "testing"

func TestMeasurePureLines_Long(t *testing.T) {
	reports, err := MeasurePureLines("long.go", longLoopSrc)
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	if len(reports) != 1 {
		t.Fatalf("want 1 report, got %d", len(reports))
	}
	if reports[0].PureLines <= 10 {
		t.Fatalf("want pure>10, got %d", reports[0].PureLines)
	}
}
