//ff:func feature=gen-ir type=test control=sequence
//ff:what TestConvert* — direct branch coverage for the per-sequence IR converters
package ir

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestConvertGetBranches_ZeroCov(t *testing.T) {
	// pagination args separated, list wrapper, nil result
	seq := ssac.Sequence{
		Type:   ssac.SeqGet,
		Model:  "Course.List",
		Inputs: map[string]string{"PerPage": "request.per_page", "InstructorID": "request.iid"},
		Result: &ssac.Result{Var: "courses", Type: "Course", Wrapper: "Page"},
	}
	op := convertGet(seq, &yongol.Fullstack{})
	if op.Kind != OpGet || op.Get == nil {
		t.Fatalf("expected OpGet, got %+v", op)
	}
	if len(op.Get.PaginationArgs) != 1 || op.Get.PaginationArgs[0].Key != "PerPage" {
		t.Errorf("pagination args = %+v", op.Get.PaginationArgs)
	}
	if len(op.Get.Args) != 1 || op.Get.Args[0].Key != "InstructorID" {
		t.Errorf("where args = %+v", op.Get.Args)
	}
	if !op.Get.IsList {
		t.Error("expected IsList for Page wrapper")
	}

	// []-prefixed list type
	op = convertGet(ssac.Sequence{Model: "Course.List", Result: &ssac.Result{Var: "cs", Type: "[]Course"}}, nil)
	if !op.Get.IsList {
		t.Error("expected IsList for []-prefixed type")
	}

	// count result type
	op = convertGet(ssac.Sequence{Model: "Course.Count", Result: &ssac.Result{Var: "n", Type: "int64"}}, nil)
	if !op.Get.IsCount {
		t.Error("expected IsCount for int64")
	}

	// nil result
	op = convertGet(ssac.Sequence{Model: "Course.X"}, nil)
	if op.Get.VarName != "" {
		t.Errorf("expected empty VarName for nil result, got %q", op.Get.VarName)
	}
}
