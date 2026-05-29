//ff:func feature=validate type=test control=sequence topic=ssac-sqlc
//ff:what TestXQS20_Case9 — @delete (`:exec`, no Result binding) → diag 0 (skip)

package ssac_sqlc

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	sqlcparser "github.com/park-jun-woo/yongol/pkg/parser/sqlc"
	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXQS20_Case9_DeleteExecPasses(t *testing.T) {
	fs := &yongol.Fullstack{
		DDLTables: []ddl.Table{workflowTable()},
		ServiceFuncs: []ssacparser.ServiceFunc{{
			Name: "DeleteWorkflow", FileName: "delete.ssac",
			Sequences: []ssacparser.Sequence{{
				Type:  "delete",
				Model: "Workflow.Delete",
				Line:  10,
			}},
		}},
		SQLcQueries: []sqlcparser.QuerySpec{makeQuery(
			"WorkflowDelete", "Workflow", "Delete", "exec",
			"DELETE FROM workflows WHERE id = @id;",
		)},
	}
	diags := runXQS20(fs)
	if len(diags) != 0 {
		t.Fatalf("expected 0 diags, got %d: %v", len(diags), diags)
	}
}
