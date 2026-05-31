//ff:func feature=external type=test control=sequence
//ff:what TestExtract* — path/body 파라미터·반환타입·키정렬·오퍼레이션 조회 검증
package external

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestFindOperation(t *testing.T) {
	getOp := &openapi3.Operation{OperationID: "list"}
	doc := &openapi3.T{Paths: openapi3.NewPaths()}
	doc.Paths.Set("/items", &openapi3.PathItem{Get: getOp})

	if got := findOperation(doc, "GET", "/items"); got != getOp {
		t.Errorf("findOperation(GET /items) = %v, want %v", got, getOp)
	}
	if got := findOperation(doc, "POST", "/items"); got != nil {
		t.Errorf("findOperation(POST /items) = %v, want nil", got)
	}
	if got := findOperation(doc, "GET", "/missing"); got != nil {
		t.Errorf("findOperation(GET /missing) = %v, want nil", got)
	}
}
