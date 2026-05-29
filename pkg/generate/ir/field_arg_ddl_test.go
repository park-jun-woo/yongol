//ff:func feature=gen-ir type=test control=sequence
//ff:what TestFieldArgDDL -- FieldArg.ColumnName/IsPK DDL 매핑 + @call/@auth SourceColumn 검증

package ir

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestFieldArgDDLColumnNameAndIsPK(t *testing.T) {
	fs := &yongol.Fullstack{
		DDLTables: []ddl.Table{
			{
				Name: "courses",
				Columns: map[string]ddl.Column{
					"id":    {Name: "id"},
					"title": {Name: "title"},
				},
				PrimaryKey: []string{"id"},
			},
		},
	}

	sf := &ssac.ServiceFunc{
		Name:     "GetCourse",
		FileName: "get_course.ssac",
		Sequences: []ssac.Sequence{
			{
				Type:  ssac.SeqGet,
				Model: "Course.FindByID",
				Inputs: map[string]string{
					"ID":    "request.id",
					"Title": "request.title",
				},
				Result: &ssac.Result{Var: "course", Type: "Course"},
			},
		},
	}

	plan, err := BuildServicePlan(sf, fs)
	if err != nil {
		t.Fatalf("BuildServicePlan: %v", err)
	}

	getOp := plan.Ops[0]
	if getOp.Kind != OpGet {
		t.Fatalf("Ops[0].Kind = %d, want OpGet", getOp.Kind)
	}

	idArg := findArgByKey(getOp.Get.Args, "ID")
	if idArg == nil {
		t.Fatal("missing ID arg")
	}
	if idArg.ColumnName != "id" {
		t.Errorf("ID.ColumnName = %q, want %q", idArg.ColumnName, "id")
	}
	if !idArg.IsPK {
		t.Error("ID.IsPK = false, want true")
	}

	titleArg := findArgByKey(getOp.Get.Args, "Title")
	if titleArg == nil {
		t.Fatal("missing Title arg")
	}
	if titleArg.ColumnName != "title" {
		t.Errorf("Title.ColumnName = %q, want %q", titleArg.ColumnName, "title")
	}
	if titleArg.IsPK {
		t.Error("Title.IsPK = true, want false")
	}
}

func TestFieldArgDDLNoMatch(t *testing.T) {
	fs := &yongol.Fullstack{
		DDLTables: []ddl.Table{
			{
				Name: "courses",
				Columns: map[string]ddl.Column{
					"id": {Name: "id"},
				},
				PrimaryKey: []string{"id"},
			},
		},
	}

	sf := &ssac.ServiceFunc{
		Name:     "GetCourse",
		FileName: "get_course.ssac",
		Sequences: []ssac.Sequence{
			{
				Type:  ssac.SeqGet,
				Model: "Course.FindByID",
				Inputs: map[string]string{
					"NonExistentCol": "request.foo",
				},
				Result: &ssac.Result{Var: "course", Type: "Course"},
			},
		},
	}

	plan, err := BuildServicePlan(sf, fs)
	if err != nil {
		t.Fatalf("BuildServicePlan: %v", err)
	}

	arg := findArgByKey(plan.Ops[0].Get.Args, "NonExistentCol")
	if arg == nil {
		t.Fatal("missing NonExistentCol arg")
	}
	if arg.ColumnName != "" {
		t.Errorf("ColumnName = %q, want empty", arg.ColumnName)
	}
	if arg.IsPK {
		t.Error("IsPK = true, want false for unmatched column")
	}
}

// TestFieldArgDDLCallOpSourceColumn verifies that @call ops get SourceColumn
// populated even without DDL table matching (Pass 1 enrichment).
func TestFieldArgDDLCallOpSourceColumn(t *testing.T) {
	sf := &ssac.ServiceFunc{
		Name:     "Login",
		FileName: "login.ssac",
		Sequences: []ssac.Sequence{
			{
				Type:   ssac.SeqCall,
				Model:  "auth.IssueToken",
				Inputs: map[string]string{"Email": "user.Email"},
				Result: &ssac.Result{Var: "token", Type: "TokenPair"},
			},
		},
	}

	plan, err := BuildServicePlan(sf, &yongol.Fullstack{})
	if err != nil {
		t.Fatalf("BuildServicePlan: %v", err)
	}

	callOp := plan.Ops[0]
	if callOp.Kind != OpCall {
		t.Fatalf("Ops[0].Kind = %d, want OpCall", callOp.Kind)
	}

	emailArg := findArgByKey(callOp.Call.Args, "Email")
	if emailArg == nil {
		t.Fatal("missing Email arg")
	}
	if emailArg.SourceColumn != "email" {
		t.Errorf("Email.SourceColumn = %q, want %q", emailArg.SourceColumn, "email")
	}
	if emailArg.Source != "user" {
		t.Errorf("Email.Source = %q, want %q", emailArg.Source, "user")
	}
}

// TestFieldArgDDLAuthOpSourceColumn verifies that @auth ops get SourceColumn
// populated via Pass 1 (DDL-independent enrichment).
func TestFieldArgDDLAuthOpSourceColumn(t *testing.T) {
	sf := &ssac.ServiceFunc{
		Name:     "DeleteWorkflow",
		FileName: "delete_workflow.ssac",
		Sequences: []ssac.Sequence{
			{
				Type:   ssac.SeqGet,
				Model:  "Workflow.FindByID",
				Inputs: map[string]string{"ID": "request.ID"},
				Result: &ssac.Result{Var: "wf", Type: "Workflow"},
			},
			{
				Type:   ssac.SeqAuth,
				Model:  "workflow.delete",
				Inputs: map[string]string{"ResourceID": "wf.OrgID"},
			},
		},
	}

	plan, err := BuildServicePlan(sf, &yongol.Fullstack{})
	if err != nil {
		t.Fatalf("BuildServicePlan: %v", err)
	}

	authOp := plan.Ops[1]
	if authOp.Kind != OpAuth {
		t.Fatalf("Ops[1].Kind = %d, want OpAuth", authOp.Kind)
	}

	resArg := findArgByKey(authOp.Auth.Inputs, "ResourceID")
	if resArg == nil {
		t.Fatal("missing ResourceID arg")
	}
	if resArg.SourceColumn != "org_id" {
		t.Errorf("ResourceID.SourceColumn = %q, want %q", resArg.SourceColumn, "org_id")
	}
	if resArg.Source != "wf" {
		t.Errorf("ResourceID.Source = %q, want %q", resArg.Source, "wf")
	}
}

func TestFieldArgDDLNilFullstack(t *testing.T) {
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

	// Empty Fullstack (no DDL tables).
	plan, err := BuildServicePlan(sf, &yongol.Fullstack{})
	if err != nil {
		t.Fatalf("BuildServicePlan: %v", err)
	}

	arg := findArgByKey(plan.Ops[0].Get.Args, "ID")
	if arg == nil {
		t.Fatal("missing ID arg")
	}
	if arg.ColumnName != "" {
		t.Errorf("ColumnName = %q, want empty when no DDL", arg.ColumnName)
	}
}
