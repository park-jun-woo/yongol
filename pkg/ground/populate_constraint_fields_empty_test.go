//ff:func feature=rule type=test control=sequence
//ff:what populateConstraintFields — empty input 분기: 어떤 key 도 쓰지 않음

package ground

import (
	"testing"

	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
)

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
