//ff:func feature=validate type=test control=sequence topic=hurl-openapi
//ff:what Run — Hurl↔OpenAPI 교차 검증 전체 실행 통합 검증

package hurl_openapi

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestRun(t *testing.T) {
	t.Run("empty_fullstack_triggers_xoh10", func(t *testing.T) {
		fs := &yongol.Fullstack{}
		diags := Run(fs)
		// XOH-10 fires because smoke.hurl is required but not found
		found := false
		for _, d := range diags {
			if strings.Contains(d.Message, "[XOH-10]") {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("expected XOH-10 diagnostic, got %v", diags)
		}
	})

	t.Run("xoh11_skipped_when_xoh10_fires", func(t *testing.T) {
		fs := &yongol.Fullstack{}
		diags := Run(fs)
		for _, d := range diags {
			if strings.Contains(d.Message, "[XOH-11]") {
				t.Errorf("XOH-11 should be skipped when XOH-10 fires, got %q", d.Message)
			}
		}
	})

	t.Run("nil_fields_no_panic", func(t *testing.T) {
		fs := &yongol.Fullstack{
			HurlEntries: nil,
			OpenAPIDoc:  nil,
		}
		// Should not panic
		_ = Run(fs)
	})

	t.Run("xoh11_runs_when_smoke_exists", func(t *testing.T) {
		fs := &yongol.Fullstack{
			HurlFiles: []string{"specs/tests/smoke.hurl"},
		}
		diags := Run(fs)
		// XOH-10 should not fire because smoke.hurl exists
		for _, d := range diags {
			if strings.Contains(d.Message, "[XOH-10]") {
				t.Errorf("XOH-10 should not fire when smoke.hurl exists, got %q", d.Message)
			}
		}
		// XOH-11 code path is now exercised (line 24)
	})
}
