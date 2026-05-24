//ff:func feature=validate type=test control=sequence topic=ssac-openapi
//ff:what xso18ResponseFieldUsedMulti — multiple funcs 독립 처리 검증

package openapi_ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/rule"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXso18ResponseFieldUsed_Multi(t *testing.T) {
	t.Run("multiple funcs processed independently", func(t *testing.T) {
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
				{
					Name:     "GetOrder",
					FileName: "order.ssac",
					Sequences: []ssac.Sequence{
						{Type: "response", Fields: map[string]string{
							"id": "order.ID",
						}},
					},
				},
			},
		}
		g := &rule.Ground{
			Schemas: map[string][]string{
				"OpenAPI.response.GetUser":  {"id", "name"},
				"OpenAPI.response.GetOrder": {"id", "total"},
			},
		}
		fs.SetGround(g)
		diags := xso18ResponseFieldUsed(fs)
		if len(diags) != 1 {
			t.Fatalf("expected 1, got %d: %+v", len(diags), diags)
		}
		if !strings.Contains(diags[0].Message, "GetOrder") {
			t.Errorf("expected GetOrder in message, got %s", diags[0].Message)
		}
		if !strings.Contains(diags[0].Message, "total") {
			t.Errorf("expected 'total' in message, got %s", diags[0].Message)
		}
	})
}
