//ff:func feature=rule type=test control=sequence dimension=1
//ff:what populateOpenAPIConstraints — request/response 제약을 prefix별로 populateConstraintFields에 위임

package ground

import (
	"testing"

	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
)

// TestPopulateOpenAPIConstraints_Dispatch verifies both request and response
// constraint maps are forwarded with correct prefixes.
func TestPopulateOpenAPIConstraints_Dispatch(t *testing.T) {
	reqFields := map[string]map[string]oapiparser.FieldConstraint{
		"CreateUser": {
			"email": {Type: "string", Required: true, MaxLength: intPtr(255)},
			"role":  {Type: "string", Enum: []string{"admin", "user"}},
		},
	}
	respFields := map[string]map[string]oapiparser.FieldConstraint{
		"GetUser": {
			"status": {Type: "string", Enum: []string{"active"}},
		},
	}
	fs := newMinimalFullstack(
		withRequestConstraints(reqFields),
		withResponseConstraints(respFields),
	)
	g := newGround()

	populateOpenAPIConstraints(g, fs)

	// Request: required + enumFields
	if _, ok := g.Schemas["OpenAPI.request.CreateUser.required"]; !ok {
		t.Errorf("missing OpenAPI.request.CreateUser.required")
	}
	if got := g.Schemas["OpenAPI.request.CreateUser.enumFields"]; len(got) != 1 || got[0] != "role" {
		t.Errorf("enumFields = %v, want [role]", got)
	}
	if g.Types["OpenAPI.request.CreateUser.maxLength.email"] != "255" {
		t.Errorf("maxLength.email not registered")
	}

	// Response: enumFields under response.constraint.<opID>
	if got := g.Schemas["OpenAPI.response.constraint.GetUser.enumFields"]; len(got) != 1 {
		t.Errorf("response enumFields len = %v, want 1", got)
	}
}
