//ff:func feature=gen-ir type=test control=sequence
//ff:what ir 내부 함수 by-name 직접 호출 테스트 — tsma content-aware 귀속용
package ir

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestExtractOpenAPIParams_ZeroCov(t *testing.T) {
	doc := &openapi3.T{
		Paths: openapi3.NewPaths(
			openapi3.WithPath("/courses/{id}", &openapi3.PathItem{
				Get: &openapi3.Operation{
					OperationID: "GetCourse",
					Parameters: openapi3.Parameters{
						&openapi3.ParameterRef{Value: &openapi3.Parameter{
							Name: "id", In: "path", Required: true,
							Schema: openapi3.NewSchemaRef("", openapi3.NewIntegerSchema()),
						}},
						&openapi3.ParameterRef{Value: &openapi3.Parameter{
							Name: "cursor", In: "query",
							Schema: openapi3.NewSchemaRef("", openapi3.NewStringSchema()),
						}},
					},
				},
			}),
		),
	}
	plan := &ServicePlan{OperationID: "GetCourse"}
	pp, qp := extractOpenAPIParams(&yongol.Fullstack{OpenAPIDoc: doc}, "GetCourse", plan)
	if !pp["id"] || !qp["cursor"] {
		t.Errorf("params not classified: path=%v query=%v", pp, qp)
	}
	// nil fs -> empty maps, no panic
	pp2, _ := extractOpenAPIParams(nil, "x", plan)
	if len(pp2) != 0 {
		t.Errorf("nil fs should give empty path params")
	}
}
