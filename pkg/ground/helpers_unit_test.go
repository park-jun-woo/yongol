//ff:func feature=ground type=test control=sequence topic=ddl
//ff:what TestGroundHelpers — unit tests for the pure ground helper functions
package ground

import (
	"sort"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestSingularize(t *testing.T) {
	tests := []struct{ in, want string }{
		{"users", "user"},
		{"categories", "category"},  // ies → y
		{"addresses", "address"},    // sses → ss
		{"boxes", "box"},            // xes → x
		{"audit_logs", "audit_log"}, // plain s
		{"address", "address"},      // ends in ss → unchanged
		{"workflow", "workflow"},    // no plural suffix
	}
	for _, tt := range tests {
		if got := singularize(tt.in); got != tt.want {
			t.Errorf("singularize(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}

func TestCollectSchemaFieldNames(t *testing.T) {
	schema := &openapi3.Schema{
		Properties: openapi3.Schemas{
			"id":   {Value: &openapi3.Schema{}},
			"name": {Value: &openapi3.Schema{}},
		},
	}
	got := collectSchemaFieldNames(schema)
	sort.Strings(got)
	if len(got) != 2 || got[0] != "id" || got[1] != "name" {
		t.Errorf("collectSchemaFieldNames = %v, want [id name]", got)
	}
	// No properties → nil.
	if got := collectSchemaFieldNames(&openapi3.Schema{}); got != nil {
		t.Errorf("empty schema → %v, want nil", got)
	}
}

func TestIsResponseCodeRelevant(t *testing.T) {
	tests := []struct {
		code string
		want bool
	}{
		{"200", true},
		{"404", true},
		{"500", true},
		{"100", false},
		{"301", false},
	}
	for _, tt := range tests {
		if got := isResponseCodeRelevant(tt.code); got != tt.want {
			t.Errorf("isResponseCodeRelevant(%q) = %v, want %v", tt.code, got, tt.want)
		}
	}
}
