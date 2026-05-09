//ff:func feature=validate type=test control=sequence topic=ssac-statemachine
//ff:what XSM-71 test — @state input type string 호환성 검증

package ssac_statemachine

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/ground"
	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	parsessac "github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// TestXsm71_StateInputUUIDFails verifies that XSM-71 rejects @state inputs
// where a DB UUID field (pgtype.UUID) is passed as a statemachine parameter.
func TestXsm71_StateInputUUIDFails(t *testing.T) {
	fs := &yongol.Fullstack{
		DDLTables: []ddl.Table{{
			Name: "workflows",
			Columns: map[string]ddl.Column{
				"id":     {Name: "id", RawType: "UUID", NotNull: true},
				"status": {Name: "status", RawType: "TEXT", NotNull: true},
			},
			ColumnOrder: []string{"id", "status"},
		}},
		ServiceFuncs: []parsessac.ServiceFunc{{
			Name:     "ActivateWorkflow",
			FileName: "service/activate_workflow.ssac",
			Sequences: []parsessac.Sequence{
				{
					Type:   "get",
					Result: &parsessac.Result{Var: "workflow", Type: "Workflow"},
				},
				{
					Type:      "state",
					DiagramID: "workflow",
					Line:      7,
					Inputs: map[string]string{
						"ID":     "workflow.ID",
						"Status": "workflow.Status",
					},
					Transition: "ActivateWorkflow",
				},
			},
		}},
	}
	fs.SetGround(ground.Build(fs))
	diags := xsm71StateInputType(fs)
	// workflow.ID → pgtype.UUID → ERROR; workflow.Status → string → OK
	if len(diags) != 1 {
		t.Fatalf("expected 1 diagnostic, got %d (%+v)", len(diags), diags)
	}
	if !strings.Contains(diags[0].Message, "[XSM-71]") {
		t.Errorf("expected [XSM-71] in message, got %q", diags[0].Message)
	}
	if !strings.Contains(diags[0].Message, "pgtype.UUID") {
		t.Errorf("expected pgtype.UUID in message, got %q", diags[0].Message)
	}
}

// TestXsm71_StateInputStringPasses verifies that @state with all
// string-typed inputs does not trigger XSM-71.
func TestXsm71_StateInputStringPasses(t *testing.T) {
	fs := &yongol.Fullstack{
		DDLTables: []ddl.Table{{
			Name: "workflows",
			Columns: map[string]ddl.Column{
				"status": {Name: "status", RawType: "TEXT", NotNull: true},
			},
			ColumnOrder: []string{"status"},
		}},
		ServiceFuncs: []parsessac.ServiceFunc{{
			Name:     "ActivateWorkflow",
			FileName: "service/activate_workflow.ssac",
			Sequences: []parsessac.Sequence{
				{
					Type:   "get",
					Result: &parsessac.Result{Var: "wf", Type: "Workflow"},
				},
				{
					Type:      "state",
					DiagramID: "workflow",
					Line:      5,
					Inputs: map[string]string{
						"status": "wf.Status",
					},
					Transition: "ActivateWorkflow",
				},
			},
		}},
	}
	fs.SetGround(ground.Build(fs))
	diags := xsm71StateInputType(fs)
	if len(diags) != 0 {
		t.Errorf("expected 0 diagnostics, got %d (%+v)", len(diags), diags)
	}
}
