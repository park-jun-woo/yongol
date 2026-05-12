//ff:func feature=design-parse type=test control=sequence
//ff:what TestParseFile_MissingFile — 존재하지 않는 파일 읽기 에러 검증

package design

import (
	"testing"
)

func TestParseFile_MissingFile(t *testing.T) {
	_, diags := ParseFile("/nonexistent/DESIGN.md")
	if len(diags) != 1 {
		t.Fatalf("expected 1 diagnostic, got %d", len(diags))
	}
	if diags[0].Level != "ERROR" {
		t.Errorf("expected ERROR level, got %q", diags[0].Level)
	}
}
