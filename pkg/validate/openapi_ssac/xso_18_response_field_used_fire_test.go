//ff:func feature=validate type=test control=sequence topic=ssac-openapi
//ff:what xso18ResponseFieldUsedFire — all used/unused field 검증

package openapi_ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/rule"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXso18ResponseFieldUsed_Fire(t *testing.T) {
	t.Run("all fields used no diagnostic", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{
				{
					Name:     "GetUser",
					FileName: "user.ssac",
					Sequences: []ssac.Sequence{
						{Type: "response", Fields: map[string]string{
							"id":   "user.ID",
							"name": "user.Name",
						}},
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

	t.Run("unused field raises XSO-18 diagnostic", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{
				{
					Name:     "GetUser",
					FileName: "user.ssac",
					Line:     10,
					Sequences: []ssac.Sequence{
						{Type: "response", Fields: map[string]string{
							"id": "user.ID",
						}},
					},
				},
			},
		}
		g := &rule.Ground{
			Schemas: map[string][]string{
				"OpenAPI.response.GetUser": {"id", "name", "email"},
			},
		}
		fs.SetGround(g)
		diags := xso18ResponseFieldUsed(fs)
		if len(diags) != 2 {
			t.Fatalf("expected 2, got %d: %+v", len(diags), diags)
		}
		for _, d := range diags {
			if !strings.Contains(d.Message, "XSO-18") {
				t.Errorf("Message missing XSO-18: %s", d.Message)
			}
			if d.File != "user.ssac" {
				t.Errorf("expected file user.ssac, got %s", d.File)
			}
			if d.Line != 10 {
				t.Errorf("expected line 10, got %d", d.Line)
			}
		}
	})
}
