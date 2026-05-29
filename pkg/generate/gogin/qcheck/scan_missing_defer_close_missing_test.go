//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestScanMissingDeferClose_Missing — defer Close 빠지면 DF-06 보고

package qcheck

import "testing"

func TestScanMissingDeferClose_Missing(t *testing.T) {
	findings, err := ScanMissingDeferClose("bad.go", dfDeferCloseMissingSrc)
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	if len(findings) != 1 {
		t.Fatalf("want 1 finding, got %d: %+v", len(findings), findings)
	}
	if findings[0].Category != "DF-06" {
		t.Fatalf("want DF-06, got %+v", findings[0])
	}
}
