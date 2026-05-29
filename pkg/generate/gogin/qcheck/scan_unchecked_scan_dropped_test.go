//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestScanUncheckedScan_Dropped — `_ = r.Scan(...)` 는 DF-02 보고

package qcheck

import "testing"

func TestScanUncheckedScan_Dropped(t *testing.T) {
	findings, err := ScanUncheckedScan("bad.go", dfScanDroppedSrc)
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	if len(findings) != 1 {
		t.Fatalf("want 1 finding, got %d: %+v", len(findings), findings)
	}
	if findings[0].Category != "DF-02" {
		t.Fatalf("want DF-02, got %+v", findings[0])
	}
}
