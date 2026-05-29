//ff:func feature=validate type=test control=iteration dimension=1 topic=func-check
//ff:what resolveInputType var.Field test — @get 결과 변수 필드 접근의 DDL Go 타입 해석

package ssac_func

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/rule"
)

// TestResolveInputType_VarField verifies that resolveInputType resolves
// var.Field expressions through SSaC.var → DDL.field lookup chain.
func TestResolveInputType_VarField(t *testing.T) {
	g := &rule.Ground{
		Types: map[string]string{
			"SSaC.var.ActivateWorkflow.workflow": "Workflow",
			"SSaC.var.ActivateWorkflow.wf":      "Workflow",
			"SSaC.var.ListWorkflows.wfs":         "[]Workflow",
			"SSaC.var.GetUser.u":                 "*User",
			"DDL.field.Workflow.ID":              "pgtype.UUID",
			"DDL.field.Workflow.Status":          "string",
			"DDL.field.Workflow.OrgID":           "pgtype.UUID",
			"DDL.field.User.Email":               "string",
			"DDL.field.User.ID":                  "int64",
		},
	}
	tests := []struct {
		funcName string
		value    string
		want     string
	}{
		{"ActivateWorkflow", "workflow.ID", "pgtype.UUID"},
		{"ActivateWorkflow", "workflow.Status", "string"},
		{"ActivateWorkflow", "wf.OrgID", "pgtype.UUID"},
		{"ListWorkflows", "wfs.ID", "pgtype.UUID"},       // []Workflow → Workflow
		{"GetUser", "u.Email", "string"},                  // *User → User
		{"GetUser", "u.ID", "int64"},
		{"ActivateWorkflow", "unknown.ID", ""},            // unknown var
		{"ActivateWorkflow", "workflow.Unknown", ""},      // unknown field
	}
	for _, tt := range tests {
		got := resolveInputType(g, tt.funcName, tt.value)
		if got != tt.want {
			t.Errorf("resolveInputType(g, %q, %q) = %q, want %q",
				tt.funcName, tt.value, got, tt.want)
		}
	}
}
