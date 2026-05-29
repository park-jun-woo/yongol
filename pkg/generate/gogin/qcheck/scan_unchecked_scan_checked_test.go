//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestScanUncheckedScan_Checked — sqlc 형태 err := row.Scan(...) + guard 는 0건

package qcheck

import "testing"

func TestScanUncheckedScan_Checked(t *testing.T) {
	findings, err := ScanUncheckedScan("checked.go", dfScanCheckedSrc)
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	if len(findings) != 0 {
		t.Fatalf("want 0 findings, got %d: %+v", len(findings), findings)
	}
}
