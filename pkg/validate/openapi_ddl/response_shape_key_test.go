//ff:func feature=validate type=test control=sequence topic=openapi-ddl
//ff:what responseShapeKey — $ref 는 ref:Name, inline 은 정렬된 키로 inline:.. 시그니처 생성

package openapi_ddl

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestResponseShapeKey(t *testing.T) {
	if got := responseShapeKey(compRef("Rule", "id", "name")); got != "ref:Rule" {
		t.Errorf("compRef shape = %q, want ref:Rule", got)
	}
	if got := responseShapeKey(inlineRef("name", "id")); got != "inline:id,name" {
		t.Errorf("inlineRef shape = %q, want inline:id,name (sorted)", got)
	}
	// bare ref without slash
	if got := responseShapeKey(&openapi3.SchemaRef{Ref: "Rule"}); got != "ref:Rule" {
		t.Errorf("bare ref shape = %q, want ref:Rule", got)
	}
}
