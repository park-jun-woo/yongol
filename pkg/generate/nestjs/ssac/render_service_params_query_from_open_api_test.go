//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestRenderServiceParams_QueryFromOpenAPI
package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestRenderServiceParams_QueryFromOpenAPI(t *testing.T) {
	plan := &ir.ServicePlan{
		HTTPMethod:  "GET",
		QueryParams: []ir.QueryParamMeta{{Name: "page"}},
	}
	got := renderServiceParams(plan)
	if !strings.Contains(got, "query: any") {
		t.Errorf("GET with QueryParams should include query, got %q", got)
	}
}
