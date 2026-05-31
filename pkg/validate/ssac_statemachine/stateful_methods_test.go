//ff:func feature=validate type=test control=iteration dimension=1 topic=states
//ff:what TestStatefulMethods — statefulMethods POST/PUT/DELETE tuple 반환 검증
package ssac_statemachine

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestStatefulMethods(t *testing.T) {
	post := &openapi3.Operation{OperationID: "p"}
	put := &openapi3.Operation{OperationID: "u"}
	del := &openapi3.Operation{OperationID: "d"}
	item := &openapi3.PathItem{Post: post, Put: put, Delete: del}

	ops := statefulMethods(item)
	if len(ops) != 3 {
		t.Fatalf("expected 3 ops, got %d", len(ops))
	}
	want := []struct {
		method string
		op     *openapi3.Operation
	}{
		{"POST", post}, {"PUT", put}, {"DELETE", del},
	}
	for i, w := range want {
		if ops[i].method != w.method || ops[i].op != w.op {
			t.Errorf("ops[%d] = {%q,%v}, want {%q,%v}", i, ops[i].method, ops[i].op, w.method, w.op)
		}
	}
}
