//ff:func feature=validate type=test-helper control=sequence
//ff:what subtestTestXos67ResponseFieldTypeMatchUuidFieldBoundToFuncOpenapiTypesUUIDPasses — uuid field bound to func openapi_types.UUID passes 서브테스트
package openapi_ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/rule"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func subtestTestXos67ResponseFieldTypeMatchUuidFieldBoundToFuncOpenapiTypesUUIDPasses(t *testing.T) {

	fs := &yongol.Fullstack{
		ServiceFuncs: []ssac.ServiceFunc{
			{
				Name: "cancelMatch",
				Sequences: []ssac.Sequence{
					{Type: "put", Result: &ssac.Result{Var: "result", Type: "CancelMatchResponse"}},
					{Type: "response", Fields: map[string]string{"match_id": "result.ID"}},
				},
			},
		},
	}
	g := &rule.Ground{
		Types: map[string]string{
			"OpenAPI.response.cancelMatch.match_id": "openapi_types.UUID",
			"SSaC.var.cancelMatch.result":           "CancelMatchResponse",
			// func struct fields register raw declared type (no apifield key).
			"Struct.CancelMatchResponse.ID": "openapi_types.UUID",
		},
	}
	fs.SetGround(g)
	diags := xos67ResponseFieldType(fs)
	if len(diags) != 0 {
		t.Fatalf("expected 0, got %d: %+v", len(diags), diags)
	}

}
