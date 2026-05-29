//ff:func feature=validate type=test control=sequence topic=func-check
//ff:what TestXfs70_AuthInputUUIDFails — @auth input UUID 타입 비호환 XFS-70 발생 검증

package ssac_func

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/ground"
	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	parsessac "github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// TestXfs70_AuthInputUUIDFails verifies that XFS-70 rejects @auth inputs
// where a DB UUID field (pgtype.UUID) is passed as ResourceID (string).
func TestXfs70_AuthInputUUIDFails(t *testing.T) {
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
					Type:     "auth",
					Action:   "ActivateWorkflow",
					Resource: "workflow",
					Line:     5,
					Inputs: map[string]string{
						"ResourceID": "workflow.ID",
					},
				},
			},
		}},
	}
	fs.SetGround(ground.Build(fs))
	diags := xfs70AuthInputType(fs)
	if len(diags) != 1 {
		t.Fatalf("expected 1 diagnostic, got %d (%+v)", len(diags), diags)
	}
	if !strings.Contains(diags[0].Message, "[XFS-70]") {
		t.Errorf("expected [XFS-70] in message, got %q", diags[0].Message)
	}
	if !strings.Contains(diags[0].Message, "pgtype.UUID") {
		t.Errorf("expected pgtype.UUID in message, got %q", diags[0].Message)
	}
}
