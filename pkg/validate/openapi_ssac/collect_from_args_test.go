//ff:func feature=validate type=test control=sequence topic=ssac-openapi
//ff:what collectFromArgs — request source 수집/비request 무시 검증

package openapi_ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestCollectFromArgs(t *testing.T) {
	t.Run("empty args does nothing", func(t *testing.T) {
		fields := map[string]bool{}
		collectFromArgs(fields, nil)
		if len(fields) != 0 {
			t.Errorf("expected empty, got %v", fields)
		}
	})

	t.Run("collects request fields", func(t *testing.T) {
		fields := map[string]bool{}
		args := []ssac.Arg{
			{Source: "request", Field: "CourseID"},
			{Source: "request", Field: "Name"},
			{Source: "course", Field: "ID"}, // non-request, skipped
			{Source: "request", Field: ""},  // empty field, skipped
		}
		collectFromArgs(fields, args)
		if len(fields) != 2 {
			t.Fatalf("expected 2, got %d: %v", len(fields), fields)
		}
		if !fields["CourseID"] || !fields["Name"] {
			t.Errorf("missing expected fields: %v", fields)
		}
	})
}
