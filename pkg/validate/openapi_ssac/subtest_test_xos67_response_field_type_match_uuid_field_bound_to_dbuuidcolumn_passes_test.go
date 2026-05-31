//ff:func feature=validate type=test-helper control=sequence
//ff:what subtestTestXos67ResponseFieldTypeMatchUuidFieldBoundToDBUUIDColumnPasses — uuid field bound to DB UUID column passes 서브테스트
package openapi_ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/rule"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func subtestTestXos67ResponseFieldTypeMatchUuidFieldBoundToDBUUIDColumnPasses(t *testing.T) {

	fs := &yongol.Fullstack{
		ServiceFuncs: []ssac.ServiceFunc{
			{
				Name: "getWorkflow",
				Sequences: []ssac.Sequence{
					{Type: "get", Result: &ssac.Result{Var: "wf", Type: "Workflow"}},
					{Type: "response", Fields: map[string]string{"id": "wf.ID"}},
				},
			},
		},
	}
	g := &rule.Ground{
		Types: map[string]string{
			// expected: format:uuid → openapi_types.UUID
			"OpenAPI.response.getWorkflow.id": "openapi_types.UUID",
			"SSaC.var.getWorkflow.wf":         "Workflow",
			// actual: DB UUID column resolved via api-surface type
			"DDL.apifield.Workflow.ID": "openapi_types.UUID",
			// coarse GoTypeOf projection (collapses UUID→string) — must be
			// overridden by the apifield key above.
			"Struct.Workflow.ID": "string",
		},
	}
	fs.SetGround(g)
	diags := xos67ResponseFieldType(fs)
	if len(diags) != 0 {
		t.Fatalf("expected 0, got %d: %+v", len(diags), diags)
	}

}
