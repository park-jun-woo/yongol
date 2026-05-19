//ff:func feature=validate type=test control=sequence topic=features-openapi
//ff:what XFO-01 — features op이 OpenAPI operationId에 없을 때 ERROR 진단 테스트

package features_openapi

import (
	"strings"
	"testing"

	featparser "github.com/park-jun-woo/yongol/pkg/parser/features"
)

func TestXFO01_OpNotInOpenAPI_Fires(t *testing.T) {
	fs := buildFSForXFO01(
		[]string{"GetWorkflow"},
		[]featparser.Feature{
			{Op: "CreateWorkflow", Path: "POST /workflows", Desc: "Create", Line: 2},
		},
	)
	diags := xfo01OpNotInOpenAPI(fs)
	if len(diags) != 1 {
		t.Fatalf("want 1 diag, got %d", len(diags))
	}
	if !strings.Contains(diags[0].Message, "[XFO-01]") {
		t.Errorf("want [XFO-01] prefix, got %s", diags[0].Message)
	}
	if !strings.Contains(diags[0].Message, "CreateWorkflow") {
		t.Errorf("want op name in message, got %s", diags[0].Message)
	}
}
