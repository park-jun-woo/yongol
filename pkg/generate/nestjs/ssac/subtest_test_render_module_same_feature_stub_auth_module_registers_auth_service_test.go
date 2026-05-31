//ff:func feature=gen-nestjs type=test-helper control=sequence
//ff:what subtestTestRenderModuleSameFeatureStubAuthModuleRegistersAuthService — AuthModuleRegistersAuthService 서브테스트
package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func subtestTestRenderModuleSameFeatureStubAuthModuleRegistersAuthService(t *testing.T) {

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
			},
		},
	}
	got, err := RenderModule("auth", plans)
	if err != nil {
		t.Fatalf("RenderModule: %v", err)
	}

	// AuthService must be in providers.
	if !strings.Contains(got, "AuthService,") {
		t.Errorf("expected AuthService in providers, got:\n%s", got)
	}

	// AuthService must be in exports.
	exportsIdx := strings.Index(got, "exports: [")
	if exportsIdx < 0 {
		t.Fatal("expected exports array")
	}
	exportsSection := got[exportsIdx:]
	if !strings.Contains(exportsSection, "AuthService,") {
		t.Errorf("expected AuthService in exports, got:\n%s", exportsSection)
	}

	// AuthService import statement.
	if !strings.Contains(got, "import { AuthService } from './auth.service'") {
		t.Errorf("expected AuthService import, got:\n%s", got)
	}

}
