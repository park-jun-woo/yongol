//ff:func feature=validate type=test control=iteration dimension=1 topic=openapi-ssac
//ff:what inferResponseValueType — empty/literal/bare var/dotted var 타입 추론 검증

package openapi_ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/rule"
)

func TestInferResponseValueType(t *testing.T) {
	g := &rule.Ground{
		Types: map[string]string{
			"SSaC.var.getUser.user":  "User",
			"Struct.User.Name":       "string",
			"Struct.User.ID":         "int64",
			"SSaC.var.getUser.count": "int64",
		},
	}

	tests := []struct {
		name     string
		funcName string
		value    string
		want     string
	}{
		{"empty value returns empty", "getUser", "", ""},
		{"literal string", "getUser", `"hello"`, "string"},
		{"literal int", "getUser", "42", "int64"},
		{"bare variable", "getUser", "count", "int64"},
		{"dotted variable field", "getUser", "user.Name", "string"},
		{"dotted variable int field", "getUser", "user.ID", "int64"},
		{"unknown bare var returns empty", "getUser", "unknown", ""},
		{"unknown dotted var returns empty", "getUser", "unknown.Field", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inferResponseValueType(g, tt.funcName, tt.value)
			if got != tt.want {
				t.Errorf("inferResponseValueType(%q, %q) = %q, want %q", tt.funcName, tt.value, got, tt.want)
			}
		})
	}
}

// TestInferResponseValueType_ApifieldPreference verifies the Phase018 change:
// a dotted DB-row field prefers DDL.apifield.<M>.<f> (api-surface type, e.g.
// a UUID column's openapi_types.UUID) over the coarse Struct.<M>.<f>
// (GoTypeOf=string), and falls back to Struct.* when no apifield key exists.
func TestInferResponseValueType_ApifieldPreference(t *testing.T) {
	g := &rule.Ground{
		Types: map[string]string{
			"SSaC.var.getWorkflow.wf": "Workflow",
			// UUID column: GoTypeOf collapses to string, apifield corrects it.
			"Struct.Workflow.ID":       "string",
			"DDL.apifield.Workflow.ID": "openapi_types.UUID",
			// non-DDL/func struct field: no apifield key → Struct fallback.
			"SSaC.var.getWorkflow.result": "FuncResp",
			"Struct.FuncResp.Token":       "string",
		},
	}

	if got := inferResponseValueType(g, "getWorkflow", "wf.ID"); got != "openapi_types.UUID" {
		t.Errorf("apifield preference: got %q, want openapi_types.UUID", got)
	}
	if got := inferResponseValueType(g, "getWorkflow", "result.Token"); got != "string" {
		t.Errorf("Struct fallback: got %q, want string", got)
	}
}
