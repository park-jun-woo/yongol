//ff:func feature=validate type=test control=sequence topic=domain-security
//ff:what TestDomainSecurityHelpers — unit tests for the pure domain_security helper functions
package domain_security

import (
	"reflect"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestContains(t *testing.T) {
	tests := []struct {
		s, sub string
		want   bool
	}{
		{"deleteworkflow", "workflow", true},
		{"abc", "", true},   // empty substr always contained
		{"a", "abc", false}, // substr longer than s
		{"hello", "ell", true},
		{"hello", "xyz", false},
		{"hello", "hello", true},
	}
	for _, tt := range tests {
		if got := contains(tt.s, tt.sub); got != tt.want {
			t.Errorf("contains(%q,%q) = %v, want %v", tt.s, tt.sub, got, tt.want)
		}
	}
}

func TestToLower(t *testing.T) {
	if got := toLower("DeleteWorkflow_123"); got != "deleteworkflow_123" {
		t.Errorf("toLower = %q", got)
	}
	if got := toLower(""); got != "" {
		t.Errorf("toLower empty = %q", got)
	}
}

func TestUniqueStrings(t *testing.T) {
	got := uniqueStrings([]string{"a", "b", "a", "c", "b"})
	want := []string{"a", "b", "c"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("uniqueStrings = %v, want %v", got, want)
	}
	if got := uniqueStrings(nil); got != nil {
		t.Errorf("nil input → %v, want nil", got)
	}
}

func TestHasEmptySecurity(t *testing.T) {
	if !hasEmptySecurity(&openapi3.Operation{Security: &openapi3.SecurityRequirements{}}) {
		t.Error("security: [] should be empty")
	}
	if hasEmptySecurity(&openapi3.Operation{}) {
		t.Error("nil security is not empty-security")
	}
	reqs := openapi3.SecurityRequirements{openapi3.SecurityRequirement{"x": {}}}
	if hasEmptySecurity(&openapi3.Operation{Security: &reqs}) {
		t.Error("non-empty security should not report empty")
	}
}

func TestHasNonEmptySecurity(t *testing.T) {
	reqs := openapi3.SecurityRequirements{openapi3.SecurityRequirement{"x": {}}}
	if !hasNonEmptySecurity(&openapi3.Operation{Security: &reqs}) {
		t.Error("non-empty security expected")
	}
	if hasNonEmptySecurity(&openapi3.Operation{}) {
		t.Error("nil security is not non-empty")
	}
	if hasNonEmptySecurity(&openapi3.Operation{Security: &openapi3.SecurityRequirements{}}) {
		t.Error("empty security is not non-empty")
	}
}

func TestHasDeleteAction(t *testing.T) {
	if !hasDeleteAction([]string{"read", "delete"}) {
		t.Error("expected delete present")
	}
	if hasDeleteAction([]string{"read", "write"}) {
		t.Error("expected no delete")
	}
	if hasDeleteAction(nil) {
		t.Error("nil should not contain delete")
	}
}

func TestHasRegoDeleteRule(t *testing.T) {
	res := map[string]struct{}{"workflow": {}}
	if !hasRegoDeleteRule("DeleteWorkflow", res) {
		t.Error("DeleteWorkflow should match workflow resource")
	}
	if hasRegoDeleteRule("DeleteUser", res) {
		t.Error("DeleteUser should not match workflow resource")
	}
	if hasRegoDeleteRule("DeleteWorkflow", map[string]struct{}{}) {
		t.Error("empty rego resources → no match")
	}
}

func TestFilterPagesByDomain(t *testing.T) {
	pages := []stml.PageSpec{
		{FileName: "admin/dashboard.stml"},
		{FileName: "public/home.stml"},
	}
	got := filterPagesByDomain(pages, "admin")
	if len(got) != 1 || got[0].FileName != "admin/dashboard.stml" {
		t.Errorf("filterPagesByDomain = %+v", got)
	}
	// empty dir → nil.
	if got := filterPagesByDomain(pages, ""); got != nil {
		t.Errorf("empty dir → %v, want nil", got)
	}
	// trailing slash tolerated.
	if got := filterPagesByDomain(pages, "admin/"); len(got) != 1 {
		t.Errorf("trailing slash dir → %+v", got)
	}
}

func TestClassifyPageDomain(t *testing.T) {
	dirs := map[string]string{
		"admin":  "admin",
		"public": "public",
	}
	if got := classifyPageDomain(stml.PageSpec{FileName: "admin/x.stml"}, dirs); got != "admin" {
		t.Errorf("got %q, want admin", got)
	}
	if got := classifyPageDomain(stml.PageSpec{FileName: "other/x.stml"}, dirs); got != "" {
		t.Errorf("unmatched should give empty, got %q", got)
	}
}

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
