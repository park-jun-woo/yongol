//ff:func feature=gen-ir type=test control=sequence
//ff:what TestServicePlanOpenAPIMeta -- ServicePlan OpenAPI 메타데이터 이식 검증 (SuccessStatus/PathParams/QueryParams/BodyFields)
package ir

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestServicePlanNoOpenAPI(t *testing.T) {
	sf := &ssac.ServiceFunc{
		Name:     "GetCourse",
		FileName: "get_course.ssac",
		Sequences: []ssac.Sequence{
			{Type: ssac.SeqGet, Model: "Course.FindByID",
				Inputs: map[string]string{"ID": "request.id"},
				Result: &ssac.Result{Var: "c", Type: "Course"}},
		},
	}

	plan, err := BuildServicePlan(sf, &yongol.Fullstack{})
	if err != nil {
		t.Fatalf("BuildServicePlan: %v", err)
	}
	if plan.SuccessStatus != 0 {
		t.Errorf("SuccessStatus = %d, want 0 (no OpenAPI)", plan.SuccessStatus)
	}
	if len(plan.PathParams) != 0 {
		t.Errorf("PathParams = %v, want empty", plan.PathParams)
	}
}
