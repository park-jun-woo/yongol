//ff:func feature=gen-ir type=test control=sequence
//ff:what TestGetOpPaginationArgs -- GetOp.PaginationArgs 분리 검증 (cursor/per_page/limit/offset)

package ir

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestGetOpPaginationArgs(t *testing.T) {
	sf := &ssac.ServiceFunc{
		Name:     "ListCourses",
		FileName: "list_courses.ssac",
		Sequences: []ssac.Sequence{
			{
				Type:  ssac.SeqGet,
				Model: "Course.List",
				Inputs: map[string]string{
					"InstructorID": "currentUser.ID",
					"cursor":       "request.cursor",
					"per_page":     "request.per_page",
				},
				Result: &ssac.Result{Var: "courses", Type: "[]Course", Wrapper: "[]"},
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

	// Where-clause args should only contain non-pagination keys.
	if len(getOp.Args) != 1 {
		t.Fatalf("len(Args) = %d, want 1 (InstructorID only)", len(getOp.Args))
	}
	if getOp.Args[0].Key != "InstructorID" {
		t.Errorf("Args[0].Key = %q, want InstructorID", getOp.Args[0].Key)
	}

	// Pagination args should contain cursor and per_page.
	if len(getOp.PaginationArgs) != 2 {
		t.Fatalf("len(PaginationArgs) = %d, want 2", len(getOp.PaginationArgs))
	}
	pagKeys := map[string]bool{}
	for _, a := range getOp.PaginationArgs {
		pagKeys[a.Key] = true
	}
	if !pagKeys["cursor"] || !pagKeys["per_page"] {
		t.Errorf("PaginationArgs keys = %v, want {cursor, per_page}", pagKeys)
	}
}

func TestGetOpNoPaginationArgs(t *testing.T) {
	sf := &ssac.ServiceFunc{
		Name:     "GetCourse",
		FileName: "get_course.ssac",
		Sequences: []ssac.Sequence{
			{
				Type:  ssac.SeqGet,
				Model: "Course.FindByID",
				Inputs: map[string]string{
					"ID": "request.id",
				},
				Result: &ssac.Result{Var: "course", Type: "Course"},
			},
		},
	}

	plan, err := BuildServicePlan(sf, &yongol.Fullstack{})
	if err != nil {
		t.Fatalf("BuildServicePlan: %v", err)
	}

	getOp := plan.Ops[0].Get
	if len(getOp.Args) != 1 {
		t.Errorf("len(Args) = %d, want 1", len(getOp.Args))
	}
	if len(getOp.PaginationArgs) != 0 {
		t.Errorf("len(PaginationArgs) = %d, want 0", len(getOp.PaginationArgs))
	}
}

// TestGetOpPaginationArgsPascalCase verifies that PascalCase pagination keys
// (e.g. "PerPage", "PageOffset") are correctly separated from where-clause
// args. This is the Group I fix.
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

func TestGetOpAllPaginationKeys(t *testing.T) {
	sf := &ssac.ServiceFunc{
		Name:     "ListItems",
		FileName: "list_items.ssac",
		Sequences: []ssac.Sequence{
			{
				Type:  ssac.SeqGet,
				Model: "Item.List",
				Inputs: map[string]string{
					"cursor":      "request.cursor",
					"per_page":    "request.per_page",
					"page_offset": "request.page_offset",
					"page":        "request.page",
					"limit":       "request.limit",
					"offset":      "request.offset",
					"Status":      "request.status",
				},
				Result: &ssac.Result{Var: "items", Type: "[]Item", Wrapper: "[]"},
			},
		},
	}

	plan, err := BuildServicePlan(sf, &yongol.Fullstack{})
	if err != nil {
		t.Fatalf("BuildServicePlan: %v", err)
	}

	getOp := plan.Ops[0].Get
	if len(getOp.Args) != 1 {
		t.Errorf("len(Args) = %d, want 1 (Status only)", len(getOp.Args))
	}
	if len(getOp.PaginationArgs) != 6 {
		t.Errorf("len(PaginationArgs) = %d, want 6", len(getOp.PaginationArgs))
	}
}
