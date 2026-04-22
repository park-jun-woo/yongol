//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestScanMissingDeferClose_Ok — boot 템플릿 형태 (defer rows.Close 포함) 는 0건

package qcheck

import "testing"

func TestScanMissingDeferClose_Ok(t *testing.T) {
	findings, err := ScanMissingDeferClose("ok.go", dfDeferCloseOkSrc)
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	if len(findings) != 0 {
		t.Fatalf("want 0 findings, got %d: %+v", len(findings), findings)
	}
}
