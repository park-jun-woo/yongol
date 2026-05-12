//ff:func feature=validate type=test control=sequence topic=ssac-statemachine
//ff:what TestXsm71_StateInputStringPasses — string 타입 @state input 은 XSM-71 미발생 검증

package ssac_statemachine

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/ground"
	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	parsessac "github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

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
