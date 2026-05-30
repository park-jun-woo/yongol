//ff:func feature=validate type=test topic=ssac-func
//ff:what zz_zerocov_test — ssac_func 0% 규칙 (Run / xfs42 / xfs43 / xfs45 / xsf46) 단위 테스트
package ssac_func

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/ground"
	"github.com/park-jun-woo/yongol/pkg/parser/funcspec"
	parsessac "github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// buildFSZeroCov wires a single @call sequence against one project FuncSpec.
func buildFSZeroCov(seq parsessac.Sequence, spec funcspec.FuncSpec) *yongol.Fullstack {
	fs := &yongol.Fullstack{
		ServiceFuncs: []parsessac.ServiceFunc{{
			Name:      "Handler",
			FileName:  "svc/handler.ssac",
			Sequences: []parsessac.Sequence{seq},
		}},
		ProjectFuncSpecs: []funcspec.FuncSpec{spec},
	}
	fs.SetGround(ground.Build(fs))
	return fs
}

func TestRun_ZeroCov(t *testing.T) {
	if diags := Run(&yongol.Fullstack{}); len(diags) != 0 {
		t.Fatalf("empty fullstack → 0 diags, got %d: %+v", len(diags), diags)
	}
}

func TestXfs42CallInputsCount_ZeroCov(t *testing.T) {
	// 2 inputs vs 1 request field → fires.
	seq := parsessac.Sequence{
		Type:   "call",
		Model:  "billing.CheckCredits",
		Line:   3,
		Inputs: map[string]string{"OrgID": "x", "Extra": "y"},
	}
	spec := funcspec.FuncSpec{
		Package:       "billing",
		Name:          "checkCredits",
		RequestFields: []funcspec.Field{{Name: "OrgID", Type: "string"}},
	}
	fs := buildFSZeroCov(seq, spec)
	if got := xfs42CallInputsCount(fs); len(got) != 1 {
		t.Fatalf("expected 1 diag, got %d: %+v", len(got), got)
	}
	// nil Ground → nil.
	if got := xfs42CallInputsCount(&yongol.Fullstack{}); got != nil {
		t.Errorf("nil ground should return nil, got %v", got)
	}
}

func TestXfs43CallInputFields_ZeroCov(t *testing.T) {
	// input key not in request fields → fires.
	seq := parsessac.Sequence{
		Type:   "call",
		Model:  "billing.CheckCredits",
		Line:   4,
		Inputs: map[string]string{"Unknown": "x"},
	}
	spec := funcspec.FuncSpec{
		Package:       "billing",
		Name:          "checkCredits",
		RequestFields: []funcspec.Field{{Name: "OrgID", Type: "string"}},
	}
	fs := buildFSZeroCov(seq, spec)
	if got := xfs43CallInputFields(fs); len(got) != 1 {
		t.Fatalf("expected 1 diag, got %d: %+v", len(got), got)
	}
	if got := xfs43CallInputFields(&yongol.Fullstack{}); got != nil {
		t.Errorf("nil ground should return nil, got %v", got)
	}
}

func TestXfs45CallResultMissing_ZeroCov(t *testing.T) {
	// @result bound but func has no Response fields → fires.
	seq := parsessac.Sequence{
		Type:   "call",
		Model:  "billing.CheckCredits",
		Line:   5,
		Result: &parsessac.Result{Var: "out"},
	}
	spec := funcspec.FuncSpec{
		Package:        "billing",
		Name:           "checkCredits",
		ResponseFields: nil,
	}
	fs := buildFSZeroCov(seq, spec)
	if got := xfs45CallResultMissing(fs); len(got) != 1 {
		t.Fatalf("expected 1 diag, got %d: %+v", len(got), got)
	}
}

func TestXsf46CallResultIgnored_ZeroCov(t *testing.T) {
	// no @result but func has Response fields → warning.
	seq := parsessac.Sequence{
		Type:   "call",
		Model:  "billing.CheckCredits",
		Line:   6,
		Result: nil,
	}
	spec := funcspec.FuncSpec{
		Package:        "billing",
		Name:           "checkCredits",
		ResponseFields: []funcspec.Field{{Name: "Balance", Type: "int"}},
	}
	fs := buildFSZeroCov(seq, spec)
	if got := xsf46CallResultIgnored(fs); len(got) != 1 {
		t.Fatalf("expected 1 diag, got %d: %+v", len(got), got)
	}
}
