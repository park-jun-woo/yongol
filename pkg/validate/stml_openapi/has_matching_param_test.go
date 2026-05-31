//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what TestStmlOpenAPIHelpers — unit tests for the pure stml_openapi helper functions
package stml_openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestHasMatchingParam(t *testing.T) {
	op := &openapi3.Operation{
		Parameters: openapi3.Parameters{
			{Value: &openapi3.Parameter{Name: "userID"}},
		},
	}
	entry := operationEntry{method: "GET", op: op}
	if !hasMatchingParam(entry, "userid") {
		t.Error("expected case-insensitive match on userID")
	}
	if hasMatchingParam(entry, "missing") {
		t.Error("missing param should not match")
	}
	// nil op → false.
	if hasMatchingParam(operationEntry{}, "x") {
		t.Error("nil op should yield no match")
	}
	// nil parameter entry / nil value → skipped (continue branch).
	opNilParam := &openapi3.Operation{
		Parameters: openapi3.Parameters{
			nil,
			{Value: nil},
			{Value: &openapi3.Parameter{Name: "ok"}},
		},
	}
	if !hasMatchingParam(operationEntry{method: "GET", op: opNilParam}, "ok") {
		t.Error("expected match on ok after skipping nil params")
	}
}
