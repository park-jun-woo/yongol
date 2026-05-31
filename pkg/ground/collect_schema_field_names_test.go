//ff:func feature=ground type=test control=sequence topic=ddl
//ff:what TestGroundHelpers — unit tests for the pure ground helper functions
package ground

import (
	"sort"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

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
