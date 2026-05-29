//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestScanUncheckedUnmarshal_Discarded — `_ = json.Unmarshal(...)` 는 DF-01 보고

package qcheck

import "testing"

func TestScanUncheckedUnmarshal_Discarded(t *testing.T) {
	findings, err := ScanUncheckedUnmarshal("bad.go", dfUnmarshalDiscardedSrc)
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	if len(findings) != 1 {
		t.Fatalf("want 1 finding, got %d: %+v", len(findings), findings)
	}
	if findings[0].Category != "DF-01" || findings[0].Detail != "json.Unmarshal" {
		t.Fatalf("want DF-01 json.Unmarshal, got %+v", findings[0])
	}
}
