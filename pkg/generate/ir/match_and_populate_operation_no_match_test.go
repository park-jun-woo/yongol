//ff:func feature=gen-ir type=test control=sequence
//ff:what TestMatchAndPopulateOperationNoMatch — TestMatchAndPopulateOperation -- operationID 매칭 시 plan 채우고 true, 미스 시 false 검증

package ir

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestMatchAndPopulateOperationNoMatch(t *testing.T) {
	pathItem := &openapi3.PathItem{
		Get: &openapi3.Operation{OperationID: "ListCourses"},
	}

	plan := &ServicePlan{}
	ok := matchAndPopulateOperation("DeleteCourse", "/courses", pathItem, plan,
		map[string]bool{}, map[string]bool{})
	if ok {
		t.Fatal("expected match=false for unmatched operationID")
	}
	if plan.HTTPMethod != "" {
		t.Errorf("plan should be untouched, HTTPMethod = %q", plan.HTTPMethod)
	}
}
