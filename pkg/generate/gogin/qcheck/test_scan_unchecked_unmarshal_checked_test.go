//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestScanUncheckedUnmarshal_Checked — 템플릿 형태 소스는 0건 보고

package qcheck

import "testing"

func TestScanUncheckedUnmarshal_Checked(t *testing.T) {
	findings, err := ScanUncheckedUnmarshal("checked.go", dfUnmarshalCheckedSrc)
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	if len(findings) != 0 {
		t.Fatalf("want 0 findings on template-shaped source, got %d: %+v", len(findings), findings)
	}
}
