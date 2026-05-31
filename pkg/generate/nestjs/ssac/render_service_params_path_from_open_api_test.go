//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestRenderServiceParams_PathFromOpenAPI
package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestRenderServiceParams_PathFromOpenAPI(t *testing.T) {
	plan := &ir.ServicePlan{
		HTTPMethod: "GET",
		PathParams: []string{"id"},
	}
	got := renderServiceParams(plan)
	if !strings.Contains(got, "params: any") {
		t.Errorf("GET with PathParams should include params, got %q", got)
	}
}
