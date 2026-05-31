//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestRenderModule_ExportsArray
package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestRenderModule_ExportsArray(t *testing.T) {
	plans := []*ir.ServicePlan{
		{OperationID: "CreateUser", HTTPMethod: "POST"},
		{OperationID: "GetUser", HTTPMethod: "GET"},
	}

	out, err := RenderModule("user", plans)
	if err != nil {
		t.Fatalf("RenderModule: %v", err)
	}

	if !strings.Contains(out, "exports: [") {
		t.Fatal("output must contain exports array")
	}
	if !strings.Contains(out, "CreateUserService,") {
		t.Error("exports must include CreateUserService")
	}
	if !strings.Contains(out, "GetUserService,") {
		t.Error("exports must include GetUserService")
	}

	// exports must come after providers
	provIdx := strings.Index(out, "providers: [")
	expIdx := strings.Index(out, "exports: [")
	if provIdx < 0 || expIdx < 0 {
		t.Fatal("both providers and exports must be present")
	}
	if expIdx < provIdx {
		t.Error("exports must come after providers")
	}
}
