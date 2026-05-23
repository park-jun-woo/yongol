//ff:func feature=validate type=test control=sequence topic=ssac-openapi
//ff:what collectRequestFields — empty sequences/args+inputs+fields 수집 검증

package openapi_ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestCollectRequestFields(t *testing.T) {
	t.Run("empty sequences returns empty", func(t *testing.T) {
		fn := ssac.ServiceFunc{}
		got := collectRequestFields(fn)
		if len(got) != 0 {
			t.Errorf("expected empty, got %v", got)
		}
	})

	t.Run("collects from args, inputs, and fields", func(t *testing.T) {
		fn := ssac.ServiceFunc{
			Sequences: []ssac.Sequence{
				{
					Args:   []ssac.Arg{{Source: "request", Field: "CourseID"}},
					Inputs: map[string]string{"status": "request.Status"},
					Fields: map[string]string{"name": "request.Name"},
				},
			},
		}
		got := collectRequestFields(fn)
		if len(got) != 3 {
			t.Fatalf("expected 3, got %d: %v", len(got), got)
		}
		if !got["CourseID"] || !got["Status"] || !got["Name"] {
			t.Errorf("missing expected fields: %v", got)
		}
	})
}
