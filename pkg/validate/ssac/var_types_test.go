//ff:func feature=validate type=test control=sequence dimension=1 topic=ssac-structural
//ff:what varTypes 단위 테스트

package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestVarTypes(t *testing.T) {
	t.Run("Collects_result_vars", func(t *testing.T) {
		fn := ssac.ServiceFunc{
			Name: "X",
			Sequences: []ssac.Sequence{
				{Type: "get", Result: &ssac.Result{Var: "course", Type: "Course"}},
				{Type: "post", Result: &ssac.Result{Var: "user", Type: "User"}},
				{Type: "put"},
			},
		}
		vt := varTypes(fn)
		if vt["course"] != "Course" {
			t.Errorf("expected course=Course, got %q", vt["course"])
		}
		if vt["user"] != "User" {
			t.Errorf("expected user=User, got %q", vt["user"])
		}
		if len(vt) != 2 {
			t.Errorf("expected 2 entries, got %d", len(vt))
		}
	})
	t.Run("Empty_func", func(t *testing.T) {
		vt := varTypes(ssac.ServiceFunc{})
		if len(vt) != 0 {
			t.Errorf("expected 0 entries, got %d", len(vt))
		}
	})
}
