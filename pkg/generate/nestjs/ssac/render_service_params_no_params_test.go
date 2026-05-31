//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestRenderServiceParams_NoParams
package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestRenderServiceParams_NoParams(t *testing.T) {
	plan := &ir.ServicePlan{HTTPMethod: "GET"}
	got := renderServiceParams(plan)
	if got != "user?: any" {
		t.Errorf("GET with no params = %q, want %q", got, "user?: any")
	}
}
