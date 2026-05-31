//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestRenderServiceParams_GETNoBody
package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestRenderServiceParams_GETNoBody(t *testing.T) {
	// GET should not have body even with BodyFields
	plan := &ir.ServicePlan{
		HTTPMethod: "GET",
		BodyFields: []ir.BodyFieldMeta{{Name: "title"}},
	}
	got := renderServiceParams(plan)
	if strings.Contains(got, "body: any") {
		t.Errorf("GET should not include body param, got %q", got)
	}
}
