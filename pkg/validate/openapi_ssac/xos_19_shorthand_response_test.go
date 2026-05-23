//ff:func feature=validate type=test control=sequence topic=ssac-openapi
//ff:what xos19ShorthandResponse — nil ground/no target/undeclared var 검증

package openapi_ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/rule"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXos19ShorthandResponse_Unit(t *testing.T) {
	t.Run("nil ground returns nil", func(t *testing.T) {
		fs := &yongol.Fullstack{}
		diags := xos19ShorthandResponse(fs)
		if len(diags) != 0 {
			t.Fatalf("expected nil, got %+v", diags)
		}
	})

	t.Run("no shorthand target skipped", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{
				{Name: "getUser", Sequences: []ssac.Sequence{{Type: "get"}}},
			},
		}
		g := &rule.Ground{Schemas: map[string][]string{}}
		fs.SetGround(g)
		diags := xos19ShorthandResponse(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d", len(diags))
		}
	})

	t.Run("undeclared var raises diagnostic", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{
				{
					Name:     "getUser",
					FileName: "user.ssac",
					Sequences: []ssac.Sequence{
						{Type: "response", Target: "undeclared"},
					},
				},
			},
		}
		g := &rule.Ground{Schemas: map[string][]string{}}
		fs.SetGround(g)
		diags := xos19ShorthandResponse(fs)
		if len(diags) != 1 {
			t.Fatalf("expected 1, got %d: %+v", len(diags), diags)
		}
		if !strings.Contains(diags[0].Message, "XOS-19") {
			t.Errorf("Message missing XOS-19: %s", diags[0].Message)
		}
	})
}
