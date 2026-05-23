//ff:func feature=validate type=test control=sequence topic=ssac-openapi
//ff:what extractResponseFieldKeys — no response/shorthand/explicit fields 검증

package openapi_ssac

import (
	"sort"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestExtractResponseFieldKeys(t *testing.T) {
	t.Run("no response sequence returns nil", func(t *testing.T) {
		fn := ssac.ServiceFunc{
			Sequences: []ssac.Sequence{{Type: "get"}},
		}
		got := extractResponseFieldKeys(fn)
		if got != nil {
			t.Errorf("expected nil, got %v", got)
		}
	})

	t.Run("shorthand response returns nil", func(t *testing.T) {
		fn := ssac.ServiceFunc{
			Sequences: []ssac.Sequence{
				{Type: "response", Target: "course"},
			},
		}
		got := extractResponseFieldKeys(fn)
		if got != nil {
			t.Errorf("expected nil, got %v", got)
		}
	})

	t.Run("explicit fields returns keys", func(t *testing.T) {
		fn := ssac.ServiceFunc{
			Sequences: []ssac.Sequence{
				{
					Type:   "response",
					Fields: map[string]string{"id": "course.ID", "name": "course.Name"},
				},
			},
		}
		got := extractResponseFieldKeys(fn)
		if len(got) != 2 {
			t.Fatalf("expected 2, got %d: %v", len(got), got)
		}
		sort.Strings(got)
		if got[0] != "id" || got[1] != "name" {
			t.Errorf("unexpected keys: %v", got)
		}
	})

	t.Run("empty fields returns nil", func(t *testing.T) {
		fn := ssac.ServiceFunc{
			Sequences: []ssac.Sequence{
				{Type: "response", Fields: map[string]string{}},
			},
		}
		got := extractResponseFieldKeys(fn)
		if got != nil {
			t.Errorf("expected nil, got %v", got)
		}
	})
}
