//ff:func feature=validate type=test control=sequence topic=hurl-openapi
//ff:what xoh09CheckFile — 한 hurl 파일에서 사용되지 않은 capture 탐지 검증

package hurl_openapi

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/hurl"
)

func TestXoh09CheckFile(t *testing.T) {
	t.Run("file_not_found_no_diag", func(t *testing.T) {
		diags := xoh09CheckFile("/nonexistent/file.hurl", nil)
		if len(diags) != 0 {
			t.Errorf("expected 0 diags, got %d", len(diags))
		}
	})

	t.Run("used_capture_no_diag", func(t *testing.T) {
		dir := t.TempDir()
		f := filepath.Join(dir, "test.hurl")
		os.WriteFile(f, []byte("POST /auth/login\nHTTP 200\n[Captures]\ntok: jsonpath \"$.token\"\nGET /api/users\nAuthorization: Bearer {{tok}}\n"), 0o644)

		entries := []hurl.HurlEntry{{
			Captures: []hurl.HurlCapture{{Var: "tok", Line: 4}},
		}}
		diags := xoh09CheckFile(f, entries)
		if len(diags) != 0 {
			t.Errorf("expected 0 diags, got %d: %v", len(diags), diags)
		}
	})

	t.Run("unused_capture_produces_warning", func(t *testing.T) {
		dir := t.TempDir()
		f := filepath.Join(dir, "test.hurl")
		os.WriteFile(f, []byte("POST /auth/login\nHTTP 200\n[Captures]\ntok: jsonpath \"$.token\"\n"), 0o644)

		entries := []hurl.HurlEntry{{
			Captures: []hurl.HurlCapture{{Var: "tok", Line: 4}},
		}}
		diags := xoh09CheckFile(f, entries)
		if len(diags) != 1 {
			t.Fatalf("expected 1 diag, got %d: %v", len(diags), diags)
		}
		if !strings.Contains(diags[0].Message, "[XOH-09]") {
			t.Errorf("expected [XOH-09], got %q", diags[0].Message)
		}
	})
}
