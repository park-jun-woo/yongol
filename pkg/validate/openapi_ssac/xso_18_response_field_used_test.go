//ff:func feature=validate type=test control=sequence topic=ssac-openapi
//ff:what xso18ResponseFieldUsed — nil ground/no response/shorthand/schema skip 검증

package openapi_ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/rule"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXso18ResponseFieldUsed(t *testing.T) {
	t.Run("nil ground returns nil", func(t *testing.T) {
		fs := &yongol.Fullstack{}
		diags := xso18ResponseFieldUsed(fs)
		if len(diags) != 0 {
			t.Fatalf("expected nil, got %+v", diags)
		}
	})

	t.Run("no service funcs returns nil", func(t *testing.T) {
		fs := &yongol.Fullstack{}
		g := &rule.Ground{
			Schemas: map[string][]string{},
		}
		fs.SetGround(g)
		diags := xso18ResponseFieldUsed(fs)
		if len(diags) != 0 {
			t.Fatalf("expected nil, got %+v", diags)
		}
	})

	t.Run("func without response sequence skipped", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{
				{
					Name:     "GetUser",
					FileName: "user.ssac",
					Sequences: []ssac.Sequence{
						{Type: "get"},
					},
				},
			},
		}
		g := &rule.Ground{
			Schemas: map[string][]string{
				"OpenAPI.response.GetUser": {"id", "name"},
			},
		}
		fs.SetGround(g)
		diags := xso18ResponseFieldUsed(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d: %+v", len(diags), diags)
		}
	})

	t.Run("shorthand response (Target set) skipped", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{
				{
					Name:     "GetUser",
					FileName: "user.ssac",
					Sequences: []ssac.Sequence{
						{Type: "response", Target: "user"},
					},
				},
			},
		}
		g := &rule.Ground{
			Schemas: map[string][]string{
				"OpenAPI.response.GetUser": {"id", "name"},
			},
		}
		fs.SetGround(g)
		diags := xso18ResponseFieldUsed(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d: %+v", len(diags), diags)
		}
	})

	t.Run("no OpenAPI schema for func skipped", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{
				{
					Name:     "GetUser",
					FileName: "user.ssac",
					Sequences: []ssac.Sequence{
						{Type: "response", Fields: map[string]string{"id": "user.ID"}},
					},
				},
			},
		}
		g := &rule.Ground{
			Schemas: map[string][]string{},
		}
		fs.SetGround(g)
		diags := xso18ResponseFieldUsed(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d: %+v", len(diags), diags)
		}
	})
}
