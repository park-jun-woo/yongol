//ff:func feature=gen-ir type=test control=sequence
//ff:what TestFieldArgDDL -- FieldArg.ColumnName/IsPK DDL 매핑 + @call/@auth SourceColumn 검증
package ir

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

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
