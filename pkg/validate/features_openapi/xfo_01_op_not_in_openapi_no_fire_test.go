//ff:func feature=validate type=test control=sequence topic=features-openapi
//ff:what XFO-01 — features op이 OpenAPI에 모두 있을 때 정상 통과 테스트

package features_openapi

import (
	"testing"

	featparser "github.com/park-jun-woo/yongol/pkg/parser/features"
)

func TestXFO01_OpNotInOpenAPI_NoFire(t *testing.T) {
	fs := buildFSForXFO01(
		[]string{"CreateWorkflow", "GetWorkflow"},
		[]featparser.Feature{
			{Op: "CreateWorkflow", Path: "POST /workflows", Desc: "Create", Line: 2},
			{Op: "GetWorkflow", Path: "GET /workflows/{id}", Desc: "Get", Line: 5},
		},
	)
	diags := xfo01OpNotInOpenAPI(fs)
	if len(diags) != 0 {
		t.Fatalf("want 0 diags, got %d", len(diags))
	}
}
