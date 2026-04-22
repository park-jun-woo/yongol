//ff:func feature=rule type=test control=sequence dimension=1
//ff:what populateConstraintFields non-enum/required/maxLength/format 케이스

package ground

import (
	"testing"

	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
)

// TestPopulateConstraintFields_RequiredMaxLengthFormat covers the non-enum
// branches: required → Schemas["<prefix>.required"], MaxLength → Types, Format
// → Types. When no field has Enum, Schemas["<prefix>.enumFields"] must NOT exist.
func TestPopulateConstraintFields_RequiredMaxLengthFormat(t *testing.T) {
	g := newGround()
	fields := map[string]oapiparser.FieldConstraint{
		"email": {
			Type:      "string",
			Format:    "email",
			MaxLength: intPtr(255),
			Required:  true,
		},
		"age": {
			Type:     "integer",
			Required: true,
		},
		"note": {
			Type: "string", // not required, no format
		},
	}

	populateConstraintFields(g, "OpenAPI.request.Foo", fields)

	// required names (order-agnostic)
	req, ok := g.Schemas["OpenAPI.request.Foo.required"]
	if !ok {
		t.Fatalf("required key missing")
	}
	if len(req) != 2 {
		t.Fatalf("required len = %d, want 2 (email, age); got %v", len(req), req)
	}
	saw := map[string]bool{}
	for _, n := range req {
		saw[n] = true
	}
	if !saw["email"] || !saw["age"] {
		t.Errorf("required missing entry: %v", req)
	}

	// MaxLength / Format via Types
	if got := g.Types["OpenAPI.request.Foo.maxLength.email"]; got != "255" {
		t.Errorf("maxLength.email = %q, want 255", got)
	}
	if got := g.Types["OpenAPI.request.Foo.format.email"]; got != "email" {
		t.Errorf("format.email = %q, want email", got)
	}

	// enumFields must not exist when no field has Enum.
	if _, ok := g.Schemas["OpenAPI.request.Foo.enumFields"]; ok {
		t.Errorf("enumFields should not exist when no enum fields present")
	}
}

// TestPopulateConstraintFields_Empty covers the empty-input branch — no keys
// should be written.
func TestPopulateConstraintFields_Empty(t *testing.T) {
	g := newGround()
	populateConstraintFields(g, "OpenAPI.request.Empty", map[string]oapiparser.FieldConstraint{})

	if _, ok := g.Schemas["OpenAPI.request.Empty.required"]; ok {
		t.Errorf("required should not exist for empty input")
	}
	if _, ok := g.Schemas["OpenAPI.request.Empty.enumFields"]; ok {
		t.Errorf("enumFields should not exist for empty input")
	}
}
