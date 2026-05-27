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

func TestRenderModule_EmptyFeature(t *testing.T) {
	_, err := RenderModule("", nil)
	if err == nil {
		t.Error("empty feature should return error")
	}
}

func TestRenderModule_SinglePlan(t *testing.T) {
	plans := []*ir.ServicePlan{
		{OperationID: "ListItems", HTTPMethod: "GET"},
	}

	out, err := RenderModule("item", plans)
	if err != nil {
		t.Fatalf("RenderModule: %v", err)
	}

	// Verify module class name
	if !strings.Contains(out, "export class ItemModule {}") {
		t.Error("module class should be ItemModule")
	}

	// Verify controllers
	if !strings.Contains(out, "ListItemsController,") {
		t.Error("controllers must include ListItemsController")
	}

	// Verify providers and exports match
	lines := strings.Split(out, "\n")
	providerServices := 0
	exportServices := 0
	inProviders := false
	inExports := false
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if strings.Contains(trimmed, "providers: [") {
			inProviders = true
			inExports = false
			continue
		}
		if strings.Contains(trimmed, "exports: [") {
			inExports = true
			inProviders = false
			continue
		}
		if trimmed == "]," {
			inProviders = false
			inExports = false
			continue
		}
		if inProviders && strings.HasSuffix(trimmed, "Service,") {
			providerServices++
		}
		if inExports && strings.HasSuffix(trimmed, "Service,") {
			exportServices++
		}
	}
	if providerServices != exportServices {
		t.Errorf("providers count (%d) != exports count (%d)", providerServices, exportServices)
	}
}
