//ff:func feature=generate type=test control=sequence
//ff:what TestMatchOperationByIDZeroCov — PathItem operations 에서 operationId 매칭 직접 커버

package generate

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestMatchOperationByID_ZeroCov(t *testing.T) {
	item := &openapi3.PathItem{
		Get:  &openapi3.Operation{OperationID: "GetCourse"},
		Post: &openapi3.Operation{OperationID: "CreateCourse"},
	}
	if op, ok := matchOperationByID(item, "CreateCourse"); !ok || op == nil {
		t.Errorf("CreateCourse should match")
	}
	if op, ok := matchOperationByID(item, "Nope"); ok || op != nil {
		t.Errorf("Nope should not match")
	}
}
