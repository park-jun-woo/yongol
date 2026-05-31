//ff:func feature=validate type=test control=sequence topic=openapi-ssac
//ff:what inferResponseValueType — empty/literal/bare var/dotted var 타입 추론 검증
package openapi_ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/rule"
)

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
