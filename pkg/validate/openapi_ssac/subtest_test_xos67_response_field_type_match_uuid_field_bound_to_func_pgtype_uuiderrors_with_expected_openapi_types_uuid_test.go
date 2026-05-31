//ff:func feature=validate type=test-helper control=sequence
//ff:what subtestTestXos67ResponseFieldTypeMatchUuidFieldBoundToFuncPgtypeUUIDErrorsWithExpectedOpenapiTypesUUID — uuid field bound to func pgtype.UUID errors with expected openapi_types.UUID 서브테스트
package openapi_ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/rule"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func subtestTestXos67ResponseFieldTypeMatchUuidFieldBoundToFuncPgtypeUUIDErrorsWithExpectedOpenapiTypesUUID(t *testing.T) {

	fs := &yongol.Fullstack{
		ServiceFuncs: []ssac.ServiceFunc{
			{
				Name:     "cancelMatch",
				FileName: "cancel_match.ssac",
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
			// func field wrongly declared as DB/sqlc type pgtype.UUID.
			"Struct.CancelMatchResponse.ID": "pgtype.UUID",
		},
	}
	fs.SetGround(g)
	diags := xos67ResponseFieldType(fs)
	if len(diags) != 1 {
		t.Fatalf("expected 1, got %d: %+v", len(diags), diags)
	}
	if !strings.Contains(diags[0].Message, "XOS-67") {
		t.Errorf("Message missing XOS-67: %s", diags[0].Message)
	}
	if !strings.Contains(diags[0].Message, "openapi_types.UUID") {
		t.Errorf("Message should mention expected openapi_types.UUID: %s", diags[0].Message)
	}
	if !strings.Contains(diags[0].Message, "pgtype.UUID") {
		t.Errorf("Message should mention actual pgtype.UUID: %s", diags[0].Message)
	}

}
