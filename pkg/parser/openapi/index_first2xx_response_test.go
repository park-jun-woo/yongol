//ff:func feature=openapi-parse type=test control=sequence
//ff:what TestOpenAPIHelpers — unit tests for the pure openapi parser helper functions
package openapi

import (
	"testing"
)

func TestIndexFirst2xxResponse(t *testing.T) {
	respsYAML := "" +
		"\"200\":\n" +
		"  content:\n" +
		"    application/json:\n" +
		"      schema:\n" +
		"        properties:\n" +
		"          id:\n" +
		"            type: integer\n" +
		"\"404\":\n" +
		"  description: not found\n"
	resps := parseNode(t, respsYAML)
	idx := &LineIndex{ResponseFields: map[string]map[string]int{}}
	indexFirst2xxResponse(resps, "GetThing", idx)
	if _, ok := idx.ResponseFields["GetThing"]["id"]; !ok {
		t.Errorf("expected id field line indexed, got %v", idx.ResponseFields["GetThing"])
	}
}
