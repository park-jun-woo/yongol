//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestRenderServiceParams_AllThree
package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestRenderServiceParams_AllThree(t *testing.T) {
	plan := &ir.ServicePlan{
		HTTPMethod:  "PUT",
		PathParams:  []string{"id"},
		BodyFields:  []ir.BodyFieldMeta{{Name: "title"}},
		QueryParams: []ir.QueryParamMeta{{Name: "v"}},
	}
	got := renderServiceParams(plan)

	// Order: params, body, query, user
	paramsIdx := strings.Index(got, "params: any")
	bodyIdx := strings.Index(got, "body: any")
	queryIdx := strings.Index(got, "query: any")
	userIdx := strings.Index(got, "user?: any")

	if paramsIdx < 0 || bodyIdx < 0 || queryIdx < 0 || userIdx < 0 {
		t.Fatalf("missing param in %q", got)
	}
	if !(paramsIdx < bodyIdx && bodyIdx < queryIdx && queryIdx < userIdx) {
		t.Errorf("wrong order in %q", got)
	}
}
