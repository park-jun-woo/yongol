//ff:func feature=validate type=test control=sequence topic=features-openapi
//ff:what Run — Features↔OpenAPI 교차 검증 정상 통과 테스트

package features_openapi

import (
	"testing"

	featparser "github.com/park-jun-woo/yongol/pkg/parser/features"
)

func TestRun_NoFire(t *testing.T) {
	fs := buildFSForXFO01(
		[]string{"CreateWorkflow", "GetWorkflow"},
		[]featparser.Feature{
			{Op: "CreateWorkflow", Path: "POST /workflows", Desc: "Create", Line: 2},
			{Op: "GetWorkflow", Path: "GET /workflows/{id}", Desc: "Get", Line: 5},
		},
	)
	diags := Run(fs)
	if len(diags) != 0 {
		t.Fatalf("want 0 diags when features and openapi match, got %d", len(diags))
	}
}
