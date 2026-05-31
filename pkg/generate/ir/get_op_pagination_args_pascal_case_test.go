//ff:func feature=gen-ir type=test control=iteration dimension=1
//ff:what TestGetOpPaginationArgs -- GetOp.PaginationArgs 분리 검증 (cursor/per_page/limit/offset)
package ir

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestGetOpPaginationArgsPascalCase(t *testing.T) {
	sf := &ssac.ServiceFunc{
		Name:     "ListAuditLogs",
		FileName: "list_audit_logs.ssac",
		Sequences: []ssac.Sequence{
			{
				Type:  ssac.SeqGet,
				Model: "AuditLog.ListByOrgID",
				Inputs: map[string]string{
					"OrgID":      "currentUser.OrgID",
					"PerPage":    "request.per_page",
					"PageOffset": "request.page_offset",
				},
				Result: &ssac.Result{Var: "logs", Type: "[]AuditLog", Wrapper: "[]"},
			},
		},
	}

	plan, err := BuildServicePlan(sf, &yongol.Fullstack{})
	if err != nil {
		t.Fatalf("BuildServicePlan: %v", err)
	}

	getOp := plan.Ops[0].Get
	if getOp == nil {
		t.Fatal("Ops[0].Get = nil")
	}

	// Where-clause should only have OrgID.
	if len(getOp.Args) != 1 {
		t.Fatalf("len(Args) = %d, want 1 (OrgID only)", len(getOp.Args))
	}
	if getOp.Args[0].Key != "OrgID" {
		t.Errorf("Args[0].Key = %q, want OrgID", getOp.Args[0].Key)
	}

	// Pagination should have PerPage and PageOffset.
	if len(getOp.PaginationArgs) != 2 {
		t.Fatalf("len(PaginationArgs) = %d, want 2", len(getOp.PaginationArgs))
	}
	pagKeys := map[string]bool{}
	for _, a := range getOp.PaginationArgs {
		pagKeys[a.Key] = true
	}
	if !pagKeys["PerPage"] || !pagKeys["PageOffset"] {
		t.Errorf("PaginationArgs keys = %v, want {PerPage, PageOffset}", pagKeys)
	}
}
