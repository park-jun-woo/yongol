//ff:func feature=manifest type=parser control=sequence
//ff:what manifest.yaml 미존재 시 diagnostic 반환 검증
package manifest

import "testing"

func TestLoad_NotFound(t *testing.T) {
	_, diags := Load("/nonexistent/dir")
	if len(diags) == 0 {
		t.Fatal("expected diagnostics for missing file")
	}
}
