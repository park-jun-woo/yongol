//ff:func feature=gen-nestjs type=test control=iteration dimension=1
//ff:what TestRenderModule_SinglePlan
package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

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
