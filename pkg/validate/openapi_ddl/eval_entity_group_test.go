//ff:func feature=validate type=test control=iteration dimension=1 topic=openapi-ddl
//ff:what evalEntityGroup — shape 발산 시 XDO-11 1건, 동일 inline 시 op 당 XDO-12, 동일 ref 시 무진단 검증

package openapi_ddl

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

func TestEvalEntityGroup(t *testing.T) {
	// >1 shape → single XDO-11 error
	diags := evalEntityGroup("Rule", []entityRepr{
		{opID: "GetRule", line: 10, shapeKey: "inline:id,name"},
		{opID: "UpdateRule", line: 20, shapeKey: "ref:Rule"},
	})
	if len(diags) != 1 || diags[0].Level != diagnostic.LevelError {
		t.Fatalf("expected 1 XDO-11 error, got %+v", diags)
	}
	if !strings.Contains(diags[0].Message, "[XDO-11]") {
		t.Errorf("missing [XDO-11] tag: %q", diags[0].Message)
	}

	// all inline same shape → XDO-12 per op (warning)
	diags = evalEntityGroup("Rule", []entityRepr{
		{opID: "GetRule", line: 10, shapeKey: "inline:id,name"},
		{opID: "MakeRule", line: 20, shapeKey: "inline:id,name"},
	})
	if len(diags) != 2 {
		t.Fatalf("expected 2 XDO-12, got %d: %+v", len(diags), diags)
	}
	for _, d := range diags {
		if d.Level != diagnostic.LevelWarning || !strings.Contains(d.Message, "[XDO-12]") {
			t.Errorf("expected XDO-12 warning, got %+v", d)
		}
	}

	// all ref same shape → no diagnostic
	diags = evalEntityGroup("Rule", []entityRepr{
		{opID: "GetRule", line: 10, shapeKey: "ref:Rule"},
		{opID: "MakeRule", line: 20, shapeKey: "ref:Rule"},
	})
	if len(diags) != 0 {
		t.Fatalf("expected 0 diagnostics, got %+v", diags)
	}
}
