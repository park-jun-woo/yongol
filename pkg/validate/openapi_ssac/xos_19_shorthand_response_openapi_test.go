//ff:func feature=validate type=test control=sequence topic=ssac-openapi
//ff:what xos19ShorthandResponse OpenAPI — declared var/missing field/wrapper skip 검증

package openapi_ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/rule"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXos19ShorthandResponse_OpenAPI(t *testing.T) {
	t.Run("declared var with fields checks against OpenAPI", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{
				{
					Name:     "getUser",
					FileName: "user.ssac",
					Sequences: []ssac.Sequence{
						{Type: "get", Result: &ssac.Result{Var: "user", Type: "User"}},
						{Type: "response", Target: "user"},
					},
				},
			},
		}
		g := &rule.Ground{
			Schemas: map[string][]string{
				"SSaC.response.getUser":    {"id", "email"},
				"OpenAPI.response.getUser": {"id", "email"},
			},
		}
		fs.SetGround(g)
		diags := xos19ShorthandResponse(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d: %+v", len(diags), diags)
		}
	})

	t.Run("declared var with missing OpenAPI field raises diagnostic", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{
				{
					Name:     "getUser",
					FileName: "user.ssac",
					Sequences: []ssac.Sequence{
						{Type: "get", Result: &ssac.Result{Var: "user", Type: "User"}},
						{Type: "response", Target: "user"},
					},
				},
			},
		}
		g := &rule.Ground{
			Schemas: map[string][]string{
				"SSaC.response.getUser":    {"id", "email", "phone"},
				"OpenAPI.response.getUser": {"id", "email"},
			},
		}
		fs.SetGround(g)
		diags := xos19ShorthandResponse(fs)
		if len(diags) != 1 {
			t.Fatalf("expected 1, got %d: %+v", len(diags), diags)
		}
	})

	t.Run("wrapper var skipped", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{
				{
					Name: "listUsers",
					Sequences: []ssac.Sequence{
						{Type: "get", Result: &ssac.Result{Var: "users", Wrapper: "Page"}},
						{Type: "response", Target: "users"},
					},
				},
			},
		}
		g := &rule.Ground{Schemas: map[string][]string{}}
		fs.SetGround(g)
		diags := xos19ShorthandResponse(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d: %+v", len(diags), diags)
		}
	})
}
