//ff:func feature=gen-nestjs type=test-helper control=sequence
//ff:what subtestTestRenderModuleSameFeatureStubMixedSameAndCrossFeature — MixedSameAndCrossFeature 서브테스트
package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func subtestTestRenderModuleSameFeatureStubMixedSameAndCrossFeature(t *testing.T) {

	plans := []*ir.ServicePlan{
		{
			OperationID: "Login",
			HTTPMethod:  "POST",
			Feature:     "auth",
			Ops: []ir.Op{
				{Kind: ir.OpCall, Call: &ir.CallOp{
					Package:       "auth",
					TargetFeature: "auth",
					Function:      "IssueToken",
				}},
				{Kind: ir.OpCall, Call: &ir.CallOp{
					Package:       "notification",
					TargetFeature: "notification",
					Function:      "SendWelcome",
				}},
			},
		},
	}
	got, err := RenderModule("auth", plans)
	if err != nil {
		t.Fatalf("RenderModule: %v", err)
	}

	// AuthService stub must be present.
	if !strings.Contains(got, "AuthService,") {
		t.Errorf("expected AuthService in providers, got:\n%s", got)
	}
	// Cross-feature module must be imported.
	if !strings.Contains(got, "NotificationModule") {
		t.Errorf("expected NotificationModule import, got:\n%s", got)
	}

}
