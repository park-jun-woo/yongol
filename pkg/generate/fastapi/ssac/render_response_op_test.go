//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestRenderResponseOp — ResponseOp → return 문 렌더링 (single/dict)

package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestRenderResponseOp(t *testing.T) {
	t.Run("Nil", func(t *testing.T) {
		var b strings.Builder
		renderResponseOp(&b, nil, "    ")
		if b.String() != "" {
			t.Errorf("expected empty, got %q", b.String())
		}
	})
	t.Run("SingleVar", func(t *testing.T) {
		var b strings.Builder
		renderResponseOp(&b, &ir.ResponseOp{SingleVar: "course"}, "    ")
		if b.String() != "    return course\n" {
			t.Errorf("got %q", b.String())
		}
	})
	t.Run("Fields", func(t *testing.T) {
		var b strings.Builder
		op := &ir.ResponseOp{Fields: []ir.ResponseField{
			{Name: "course", Source: "course"},
			{Name: "instructor_name", Source: "instructor.Name"},
		}}
		renderResponseOp(&b, op, "    ")
		out := b.String()
		if !strings.HasPrefix(out, "    return {\n") || !strings.HasSuffix(out, "    }\n") {
			t.Errorf("malformed dict: %q", out)
		}
		if !strings.Contains(out, `"course": course,`) {
			t.Errorf("missing course field: %q", out)
		}
		if !strings.Contains(out, `"instructor_name": instructor["name"],`) {
			t.Errorf("missing converted field: %q", out)
		}
	})
}
