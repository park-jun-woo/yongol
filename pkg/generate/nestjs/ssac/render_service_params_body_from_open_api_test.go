//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestRenderServiceParams_BodyFromOpenAPI
package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestRenderServiceParams_BodyFromOpenAPI(t *testing.T) {
	plan := &ir.ServicePlan{
		HTTPMethod: "POST",
		BodyFields: []ir.BodyFieldMeta{{Name: "title"}, {Name: "content"}},
	}
	got := renderServiceParams(plan)
	if !strings.Contains(got, "body: any") {
		t.Errorf("POST with BodyFields should include body param, got %q", got)
	}
}
