//ff:func feature=validate type=test control=sequence topic=features-openapi
//ff:what XOF-01 — OpenAPI operationId가 features에 없을 때 ERROR 진단 테스트

package features_openapi

import (
	"strings"
	"testing"

	featparser "github.com/park-jun-woo/yongol/pkg/parser/features"
)

func TestXOF01_OpIDNotInFeatures_Fires(t *testing.T) {
	fs := buildFSForXOF01(
		[]string{"CreateWorkflow", "DeleteWorkflow"},
		[]featparser.Feature{
			{Op: "CreateWorkflow", Path: "POST /workflows", Desc: "Create", Line: 2},
		},
		nil,
	)
	diags := xof01OpIDNotInFeatures(fs)
	if len(diags) != 1 {
		t.Fatalf("want 1 diag, got %d", len(diags))
	}
	if !strings.Contains(diags[0].Message, "[XOF-01]") {
		t.Errorf("want [XOF-01] prefix, got %s", diags[0].Message)
	}
	if !strings.Contains(diags[0].Message, "DeleteWorkflow") {
		t.Errorf("want op name in message, got %s", diags[0].Message)
	}
}
