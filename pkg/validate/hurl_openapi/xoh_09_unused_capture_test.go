//ff:func feature=validate type=test control=sequence topic=hurl-openapi
//ff:what xoh09UnusedCapture — Captures 변수가 파일 내에서 재사용되는지 통합 검증

package hurl_openapi

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/hurl"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXoh09UnusedCapture(t *testing.T) {
	t.Run("nil_fs", func(t *testing.T) {
		if diags := xoh09UnusedCapture(nil); len(diags) != 0 {
			t.Errorf("expected 0 diags, got %d", len(diags))
		}
	})

	t.Run("unused_capture_warning", func(t *testing.T) {
		dir := t.TempDir()
		f := filepath.Join(dir, "test.hurl")
		os.WriteFile(f, []byte("POST /auth/login\nHTTP 200\n[Captures]\ntok: jsonpath \"$.token\"\n"), 0o644)

		fs := &yongol.Fullstack{
			HurlEntries: []hurl.HurlEntry{{
				File:     f,
				Captures: []hurl.HurlCapture{{Var: "tok", Line: 4}},
			}},
		}
		diags := xoh09UnusedCapture(fs)
		if len(diags) != 1 {
			t.Fatalf("expected 1 diag, got %d: %v", len(diags), diags)
		}
		if !strings.Contains(diags[0].Message, "[XOH-09]") {
			t.Errorf("expected [XOH-09], got %q", diags[0].Message)
		}
	})
}
