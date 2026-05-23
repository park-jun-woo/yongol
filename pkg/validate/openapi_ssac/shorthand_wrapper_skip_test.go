//ff:func feature=validate type=test control=sequence topic=ssac-openapi
//ff:what shorthandWrapperSkip — no result/no wrapper/wrapper 존재 검증

package openapi_ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestShorthandWrapperSkip(t *testing.T) {
	t.Run("no result returns false", func(t *testing.T) {
		fn := ssac.ServiceFunc{
			Sequences: []ssac.Sequence{{Type: "get"}},
		}
		if shorthandWrapperSkip(fn, "course") {
			t.Error("expected false")
		}
	})

	t.Run("result var mismatch returns false", func(t *testing.T) {
		fn := ssac.ServiceFunc{
			Sequences: []ssac.Sequence{
				{Result: &ssac.Result{Var: "other", Wrapper: "Page"}},
			},
		}
		if shorthandWrapperSkip(fn, "course") {
			t.Error("expected false")
		}
	})

	t.Run("matching var without wrapper returns false", func(t *testing.T) {
		fn := ssac.ServiceFunc{
			Sequences: []ssac.Sequence{
				{Result: &ssac.Result{Var: "course", Wrapper: ""}},
			},
		}
		if shorthandWrapperSkip(fn, "course") {
			t.Error("expected false")
		}
	})

	t.Run("matching var with wrapper returns true", func(t *testing.T) {
		fn := ssac.ServiceFunc{
			Sequences: []ssac.Sequence{
				{Result: &ssac.Result{Var: "courses", Wrapper: "Page"}},
			},
		}
		if !shorthandWrapperSkip(fn, "courses") {
			t.Error("expected true")
		}
	})
}
