//ff:func feature=validate type=test control=sequence topic=features-openapi
//ff:what XOF-01 — LineIndex로 정확한 라인 번호 반환 테스트

package features_openapi

import (
	"testing"

	featparser "github.com/park-jun-woo/yongol/pkg/parser/features"
	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
)

func TestXOF01_OpIDNotInFeatures_WithLineIndex(t *testing.T) {
	idx := &oapiparser.LineIndex{
		Operations: map[string]int{"DeleteWorkflow": 42},
	}
	fs := buildFSForXOF01(
		[]string{"CreateWorkflow", "DeleteWorkflow"},
		[]featparser.Feature{
			{Op: "CreateWorkflow", Path: "POST /workflows", Desc: "Create", Line: 2},
		},
		idx,
	)
	diags := xof01OpIDNotInFeatures(fs)
	if len(diags) != 1 {
		t.Fatalf("want 1 diag, got %d", len(diags))
	}
	if diags[0].Line != 42 {
		t.Errorf("want line 42 from LineIndex, got %d", diags[0].Line)
	}
}
