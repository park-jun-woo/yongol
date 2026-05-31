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
