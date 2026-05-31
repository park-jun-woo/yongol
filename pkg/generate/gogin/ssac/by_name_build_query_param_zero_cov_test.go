//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestByName_ZeroCov — gogin/ssac 응답·INSERT·쿼리 렌더 헬퍼들을 이름으로 직접 호출해 커버리지 귀속
package ssac

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestByNameBuildQueryParam_ZeroCov(t *testing.T) {
	// nil ParameterRef → default string.
	if qp := buildQueryParam(nil, "ListItems"); qp.GoType != "string" {
		t.Errorf("nil param GoType = %q", qp.GoType)
	}
	// integer + format + enum.
	sch := openapi3.NewIntegerSchema()
	sch.Format = "int64"
	sch.Enum = []interface{}{int64(1), int64(2)}
	p := &openapi3.ParameterRef{Value: &openapi3.Parameter{
		Name:     "status",
		Required: true,
		Schema:   &openapi3.SchemaRef{Value: sch},
	}}
	qp := buildQueryParam(p, "ListItems")
	if !qp.IsRequired || !qp.IsEnum || qp.EnumTypeName == "" {
		t.Errorf("buildQueryParam enum = %+v", qp)
	}
	// schema nil branch.
	p2 := &openapi3.ParameterRef{Value: &openapi3.Parameter{Name: "x"}}
	if qp := buildQueryParam(p2, "Op"); qp.GoType != "string" {
		t.Errorf("no-schema GoType = %q", qp.GoType)
	}
}
