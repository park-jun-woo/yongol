//ff:func feature=validate type=test control=sequence topic=features-openapi
//ff:what XOF-01 — OpenAPI operationId가 모두 features에 있을 때 정상 통과 테스트

package features_openapi

import (
	"testing"

	featparser "github.com/park-jun-woo/yongol/pkg/parser/features"
)

func TestXOF01_OpIDNotInFeatures_NoFire(t *testing.T) {
	fs := buildFSForXOF01(
		[]string{"CreateWorkflow", "GetWorkflow"},
		[]featparser.Feature{
			{Op: "CreateWorkflow", Path: "POST /workflows", Desc: "Create", Line: 2},
			{Op: "GetWorkflow", Path: "GET /workflows/{id}", Desc: "Get", Line: 5},
		},
		nil,
	)
	diags := xof01OpIDNotInFeatures(fs)
	if len(diags) != 0 {
		t.Fatalf("want 0 diags, got %d", len(diags))
	}
}
