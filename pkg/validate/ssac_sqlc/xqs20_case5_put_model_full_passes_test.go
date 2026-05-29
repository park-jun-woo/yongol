//ff:func feature=validate type=test control=sequence topic=ssac-sqlc
//ff:what TestXQS20_Case5 — @put Workflow + RETURNING * → diag 0 (Model ↔ full UPDATE)

package ssac_sqlc

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	sqlcparser "github.com/park-jun-woo/yongol/pkg/parser/sqlc"
	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXQS20_Case5_PutModelFullPasses(t *testing.T) {
	fs := &yongol.Fullstack{
		DDLTables: []ddl.Table{workflowTable()},
		ServiceFuncs: []ssacparser.ServiceFunc{{
			Name: "UpdateWorkflow", FileName: "update.ssac",
			Sequences: []ssacparser.Sequence{makeSeq("put", "Workflow", "Workflow.UpdateMeta")},
		}},
		SQLcQueries: []sqlcparser.QuerySpec{makeQuery(
			"WorkflowUpdateMeta", "Workflow", "UpdateMeta", "one",
			"UPDATE workflows SET meta = @meta WHERE id = @id RETURNING *;",
		)},
	}
	diags := runXQS20(fs)
	if len(diags) != 0 {
		t.Fatalf("expected 0 diags, got %d: %v", len(diags), diags)
	}
}
