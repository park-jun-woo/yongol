//ff:func feature=gen-react type=test control=sequence
//ff:what TestByName_ZeroCov — react 코드젠 헬퍼들을 이름으로 직접 호출해 커버리지 귀속
package react

import (
	"strings"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestByNameEndpoints_ZeroCov(t *testing.T) {
	doc := &openapi3.T{
		Paths: openapi3.NewPaths(
			openapi3.WithPath("/api/items/{id}", &openapi3.PathItem{
				Get:  &openapi3.Operation{OperationID: "GetItem"},
				Post: &openapi3.Operation{OperationID: "CreateItem"},
			}),
		),
	}
	eps := collectEndpoints(doc)
	if len(eps) != 2 {
		t.Fatalf("collectEndpoints = %d, want 2", len(eps))
	}

	pi := &openapi3.PathItem{Get: &openapi3.Operation{OperationID: "X"}}
	got := appendOperations(nil, "/x", pi)
	if len(got) != 1 {
		t.Errorf("appendOperations = %d", len(got))
	}

	var b strings.Builder
	writeApiClientEntry(&b, eps[0], false)
	if b.Len() == 0 {
		t.Errorf("writeApiClientEntry empty")
	}
}
