//ff:func feature=gen-ir type=test control=sequence
//ff:what TestFieldArgDDL -- FieldArg.ColumnName/IsPK DDL 매핑 + @call/@auth SourceColumn 검증
package ir

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

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
