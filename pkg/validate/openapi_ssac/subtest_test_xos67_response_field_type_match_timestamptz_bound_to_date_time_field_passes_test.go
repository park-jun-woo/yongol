//ff:func feature=validate type=test-helper control=sequence
//ff:what subtestTestXos67ResponseFieldTypeMatchTimestamptzBoundToDateTimeFieldPasses — timestamptz bound to date-time field passes 서브테스트
package openapi_ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/rule"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func subtestTestXos67ResponseFieldTypeMatchTimestamptzBoundToDateTimeFieldPasses(t *testing.T) {

	fs := &yongol.Fullstack{
		ServiceFuncs: []ssac.ServiceFunc{
			{
				Name: "approveItem",
				Sequences: []ssac.Sequence{
					{Type: "put", Result: &ssac.Result{Var: "updated", Type: "Item"}},
					{Type: "response", Fields: map[string]string{"approved_at": "updated.ApprovedAt"}},
				},
			},
		},
	}
	g := &rule.Ground{
		Types: map[string]string{
			"OpenAPI.response.approveItem.approved_at": "time.Time",
			"SSaC.var.approveItem.updated":             "Item",
			"Struct.Item.ApprovedAt":                   "time.Time",
		},
	}
	fs.SetGround(g)
	diags := xos67ResponseFieldType(fs)
	if len(diags) != 0 {
		t.Fatalf("expected 0, got %d: %+v", len(diags), diags)
	}

}
