//ff:func feature=validate type=test control=sequence topic=ssac-openapi
//ff:what xso20ShorthandFieldUsedFire — all used/unused field 검증

package openapi_ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/rule"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXso20ShorthandFieldUsed_Fire(t *testing.T) {
	t.Run("all fields used no diagnostic", func(t *testing.T) {
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
				"SSaC.response.GetUser":    {"id", "name"},
				"OpenAPI.response.GetUser": {"id", "name"},
			},
		}
		fs.SetGround(g)
		diags := xso20ShorthandFieldUsed(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d: %+v", len(diags), diags)
		}
	})

	t.Run("unused field raises XSO-20 diagnostic", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{
				{
					Name:     "GetUser",
					FileName: "user.ssac",
					Line:     5,
					Sequences: []ssac.Sequence{
						{Type: "response", Target: "user"},
					},
				},
			},
		}
		g := &rule.Ground{
			Schemas: map[string][]string{
				"SSaC.response.GetUser":    {"id"},
				"OpenAPI.response.GetUser": {"id", "name", "email"},
			},
		}
		fs.SetGround(g)
		diags := xso20ShorthandFieldUsed(fs)
		if len(diags) != 2 {
			t.Fatalf("expected 2, got %d: %+v", len(diags), diags)
		}
		for _, d := range diags {
			if !strings.Contains(d.Message, "XSO-20") {
				t.Errorf("Message missing XSO-20: %s", d.Message)
			}
			if d.File != "user.ssac" {
				t.Errorf("expected file user.ssac, got %s", d.File)
			}
			if d.Line != 5 {
				t.Errorf("expected line 5, got %d", d.Line)
			}
			if !strings.Contains(d.Advice, "SSaC variable") {
				t.Errorf("expected XSO-20 specific advice, got %s", d.Advice)
			}
		}
	})
}
