//ff:func feature=validate type=test control=sequence topic=ssac-openapi
//ff:what collectMissingProps — 매칭/누락/ruleID별 advice 검증

package openapi_ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestCollectMissingProps(t *testing.T) {
	fn := ssac.ServiceFunc{Name: "getUser", FileName: "user.ssac", Line: 10}
	opProps := map[string]bool{"id": true, "email": true}

	t.Run("all fields present no diagnostics", func(t *testing.T) {
		diags := collectMissingProps(fn, []string{"id", "email"}, opProps, "XOS-17")
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d", len(diags))
		}
	})

	t.Run("missing field raises diagnostic", func(t *testing.T) {
		diags := collectMissingProps(fn, []string{"phone"}, opProps, "XOS-17")
		if len(diags) != 1 {
			t.Fatalf("expected 1, got %d", len(diags))
		}
		if !strings.Contains(diags[0].Message, "XOS-17") {
			t.Errorf("Message missing XOS-17: %s", diags[0].Message)
		}
	})

	t.Run("XOS-19 advice", func(t *testing.T) {
		diags := collectMissingProps(fn, []string{"phone"}, opProps, "XOS-19")
		if len(diags) != 1 {
			t.Fatalf("expected 1, got %d", len(diags))
		}
		if !strings.Contains(diags[0].Advice, "shorthand") {
			t.Errorf("Advice missing shorthand: %s", diags[0].Advice)
		}
	})

	t.Run("XSO-20 advice", func(t *testing.T) {
		diags := collectMissingProps(fn, []string{"phone"}, opProps, "XSO-20")
		if len(diags) != 1 {
			t.Fatalf("expected 1, got %d", len(diags))
		}
		if !strings.Contains(diags[0].Advice, "remove it from @response") {
			t.Errorf("Advice missing expected text: %s", diags[0].Advice)
		}
	})

	t.Run("unknown ruleID default advice", func(t *testing.T) {
		diags := collectMissingProps(fn, []string{"phone"}, opProps, "UNKNOWN")
		if len(diags) != 1 {
			t.Fatalf("expected 1, got %d", len(diags))
		}
		if !strings.Contains(diags[0].Advice, "Add @response") {
			t.Errorf("Advice missing default text: %s", diags[0].Advice)
		}
	})
}
