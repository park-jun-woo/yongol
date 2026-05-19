//ff:func feature=validate type=test control=sequence topic=features-openapi
//ff:what Run — Features↔OpenAPI 교차 검증 두 규칙 동시 발동 테스트

package features_openapi

import (
	"testing"

	featparser "github.com/park-jun-woo/yongol/pkg/parser/features"
)

func TestRun_BothRulesFire(t *testing.T) {
	// "CreateWorkflow" is in features but not in OpenAPI (XFO-01).
	// "DeleteWorkflow" is in OpenAPI but not in features (XOF-01).
	fs := buildFSForXFO01(
		[]string{"DeleteWorkflow"},
		[]featparser.Feature{
			{Op: "CreateWorkflow", Path: "POST /workflows", Desc: "Create", Line: 2},
		},
	)
	diags := Run(fs)
	if len(diags) != 2 {
		t.Fatalf("want 2 diags (one XFO-01, one XOF-01), got %d", len(diags))
	}
}
