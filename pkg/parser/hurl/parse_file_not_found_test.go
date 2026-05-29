//ff:func feature=scenario type=parser control=sequence
//ff:what 존재하지 않는 파일에 대해 diagnostic 반환 검증
package hurl

import "testing"

func TestParseFile_NotFound(t *testing.T) {
	entries, diags := ParseFile("/nonexistent/file.hurl")
	if entries != nil {
		t.Fatalf("expected nil entries, got %v", entries)
	}
	if len(diags) == 0 {
		t.Fatal("expected diagnostics for missing file")
	}
}
