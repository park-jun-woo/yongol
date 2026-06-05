//ff:func feature=validate type=test control=iteration dimension=1 topic=stml-openapi
//ff:what TestOperationsOf — PathItem의 GET/POST/PUT/DELETE/PATCH operation을 고정 순서로 반환(nil 보존) 검증
package stml_openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestOperationsOf(t *testing.T) {
	get := &openapi3.Operation{OperationID: "Get"}
	post := &openapi3.Operation{OperationID: "Post"}
	put := &openapi3.Operation{OperationID: "Put"}
	del := &openapi3.Operation{OperationID: "Delete"}
	patch := &openapi3.Operation{OperationID: "Patch"}

	item := &openapi3.PathItem{Get: get, Post: post, Put: put, Delete: del, Patch: patch}
	ops := operationsOf(item)

	want := []*openapi3.Operation{get, post, put, del, patch}
	if len(ops) != len(want) {
		t.Fatalf("expected %d ops, got %d", len(want), len(ops))
	}
	for i, w := range want {
		if ops[i] != w {
			t.Errorf("position %d: got %v, want %v", i, ops[i], w)
		}
	}

	// Undefined verbs are preserved as nil entries in the same fixed order.
	partial := operationsOf(&openapi3.PathItem{Get: get})
	if len(partial) != 5 {
		t.Fatalf("expected 5 entries, got %d", len(partial))
	}
	if partial[0] != get {
		t.Errorf("position 0: want get op")
	}
	for i := 1; i < 5; i++ {
		if partial[i] != nil {
			t.Errorf("position %d: expected nil, got %v", i, partial[i])
		}
	}
}
