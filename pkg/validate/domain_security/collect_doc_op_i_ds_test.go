//ff:func feature=validate type=test control=sequence topic=domain-security
//ff:what TestDomainSecurityHelpers — unit tests for the pure domain_security helper functions
package domain_security

import (
	"reflect"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestCollectDocOpIDs(t *testing.T) {
	doc := &openapi3.T{Paths: openapi3.NewPaths()}
	doc.Paths.Set("/users", &openapi3.PathItem{
		Get:    &openapi3.Operation{OperationID: "ListUsers"},
		Post:   &openapi3.Operation{OperationID: "CreateUser"},
		Delete: &openapi3.Operation{OperationID: ""}, // empty id skipped
	})
	opDomains := map[string][]string{}
	collectDocOpIDs(domainDoc{Name: "core", Doc: doc}, opDomains)
	if !reflect.DeepEqual(opDomains["ListUsers"], []string{"core"}) {
		t.Errorf("ListUsers domains = %v", opDomains["ListUsers"])
	}
	if !reflect.DeepEqual(opDomains["CreateUser"], []string{"core"}) {
		t.Errorf("CreateUser domains = %v", opDomains["CreateUser"])
	}
	if _, ok := opDomains[""]; ok {
		t.Error("empty operationId should be skipped")
	}

	// nil Paths → no-op.
	collectDocOpIDs(domainDoc{Name: "x", Doc: &openapi3.T{}}, opDomains)
}
