//ff:func feature=validate type=test control=sequence dimension=2 topic=ssac-structural
//ff:what S-47 — CRUD Model 에 package prefix 금지 검증

package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestS47NoDotPrefix(t *testing.T) {
	t.Run("Fires_double_dot", func(t *testing.T) {
		// "pkg.Order.FindByID" has 2 dots — package prefix detected
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{
				{FileName: "order.ssac", Sequences: []ssac.Sequence{
					{Type: "get", Line: 3, Model: "pkg.Order.FindByID"},
				}},
			},
		}
		diags := s47NoDotPrefix(fs)
		found := false
		for _, d := range diags {
			if strings.Contains(d.Message, "[S-47]") {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("expected S-47 diagnostic for double-dot model")
		}
	})

	t.Run("Fires_package_field", func(t *testing.T) {
		// Package field set + 1 dot in Model
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{
				{FileName: "order.ssac", Sequences: []ssac.Sequence{
					{Type: "get", Line: 3, Model: "Order.FindByID", Package: "session"},
				}},
			},
		}
		diags := s47NoDotPrefix(fs)
		found := false
		for _, d := range diags {
			if strings.Contains(d.Message, "[S-47]") {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("expected S-47 diagnostic for Package + Model")
		}
	})

	t.Run("Passes_single_dot", func(t *testing.T) {
		// "Order.FindByID" — one dot, no package — OK
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{
				{FileName: "order.ssac", Sequences: []ssac.Sequence{
					{Type: "get", Line: 3, Model: "Order.FindByID"},
				}},
			},
		}
		diags := s47NoDotPrefix(fs)
		for _, d := range diags {
			if strings.Contains(d.Message, "[S-47]") {
				t.Errorf("unexpected S-47 diagnostic: %s", d.Message)
			}
		}
	})

	t.Run("Skips_non_crud", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{
				{FileName: "order.ssac", Sequences: []ssac.Sequence{
					{Type: "empty", Line: 3, Target: "order"},
				}},
			},
		}
		diags := s47NoDotPrefix(fs)
		if len(diags) != 0 {
			t.Errorf("got %d diags, want 0 for non-CRUD", len(diags))
		}
	})

	t.Run("Skips_empty_model", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{
				{FileName: "order.ssac", Sequences: []ssac.Sequence{
					{Type: "get", Line: 3, Model: ""},
				}},
			},
		}
		diags := s47NoDotPrefix(fs)
		if len(diags) != 0 {
			t.Errorf("got %d diags, want 0 for empty model", len(diags))
		}
	})

	t.Run("Empty", func(t *testing.T) {
		fs := &yongol.Fullstack{}
		diags := s47NoDotPrefix(fs)
		if len(diags) != 0 {
			t.Errorf("got %d diags, want 0", len(diags))
		}
	})
}
