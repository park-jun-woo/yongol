//ff:func feature=validate type=test control=sequence topic=ssac-openapi
//ff:what shorthandResponseTarget — no response/explicit/shorthand 검증

package openapi_ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestShorthandResponseTarget(t *testing.T) {
	t.Run("no response returns empty", func(t *testing.T) {
		fn := ssac.ServiceFunc{Sequences: []ssac.Sequence{{Type: "get"}}}
		if got := shorthandResponseTarget(fn); got != "" {
			t.Errorf("expected empty, got %q", got)
		}
	})

	t.Run("explicit response (no target) returns empty", func(t *testing.T) {
		fn := ssac.ServiceFunc{
			Sequences: []ssac.Sequence{
				{Type: "response", Fields: map[string]string{"id": "course.ID"}},
			},
		}
		if got := shorthandResponseTarget(fn); got != "" {
			t.Errorf("expected empty, got %q", got)
		}
	})

	t.Run("shorthand response returns target", func(t *testing.T) {
		fn := ssac.ServiceFunc{
			Sequences: []ssac.Sequence{
				{Type: "response", Target: "course"},
			},
		}
		if got := shorthandResponseTarget(fn); got != "course" {
			t.Errorf("expected course, got %q", got)
		}
	})
}
