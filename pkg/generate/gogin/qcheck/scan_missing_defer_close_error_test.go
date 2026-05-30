//ff:func feature=gen-gogin type=test control=branch topic=defensive
//ff:what TestScanMissingDeferClose_ParseError — 잘못된 소스에서 파서 에러 전파 검증

package qcheck

import "testing"

func TestScanMissingDeferClose_ParseError(t *testing.T) {
	_, err := ScanMissingDeferClose("bad.go", "@@@ not go")
	if err == nil {
		t.Fatalf("expected parse error, got nil")
	}
}
