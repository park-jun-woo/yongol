//ff:func feature=rule type=test control=sequence dimension=1
//ff:what populateConstraintFields enum 케이스 — PhaseV09 enumFields append 회귀 방지

package ground

import (
	"testing"

	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
)

// TestPopulateConstraintFields_EnumFieldsAppended guards against the dead-code
// regression fixed in PhaseV09: enumFields must be appended and registered
// under Schemas["<prefix>.enumFields"] when any field has an Enum list.
func TestPopulateConstraintFields_EnumFieldsAppended(t *testing.T) {
	g := newGround()
	fields := map[string]oapiparser.FieldConstraint{
		"status": {
			Type: "string",
			Enum: []string{"active", "inactive"},
		},
		"kind": {
			Type: "string",
			Enum: []string{"alpha", "beta", "gamma"},
		},
		"name": {
			Type: "string", // no enum
		},
	}

	populateConstraintFields(g, "OpenAPI.request.Foo", fields)

	got, ok := g.Schemas["OpenAPI.request.Foo.enumFields"]
	if !ok {
		t.Fatal("Schemas[OpenAPI.request.Foo.enumFields] missing — PhaseV09 regression")
	}
	if len(got) != 2 {
		t.Fatalf("enumFields len = %d, want 2 (status, kind); got %v", len(got), got)
	}
	seen := map[string]bool{}
	for _, n := range got {
		seen[n] = true
	}
	if !seen["status"] || !seen["kind"] {
		t.Errorf("enumFields missing expected entries; got %v", got)
	}
	if seen["name"] {
		t.Errorf("enumFields unexpectedly contains non-enum field 'name'; got %v", got)
	}

	// Schemas[...enum.<field>] must carry values.
	if vs := g.Schemas["OpenAPI.request.Foo.enum.status"]; len(vs) != 2 {
		t.Errorf("enum.status values = %v; want 2", vs)
	}
	// Types[...enum.<field>] must carry CSV.
	if g.Types["OpenAPI.request.Foo.enum.status"] == "" {
		t.Errorf("Types[enum.status] empty")
	}
}
