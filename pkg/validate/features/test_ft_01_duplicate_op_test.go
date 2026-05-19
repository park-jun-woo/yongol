//ff:what FT-01 — op 중복 시 에러 진단 테스트
package features

import (
	"strings"
	"testing"

	featparser "github.com/park-jun-woo/yongol/pkg/parser/features"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestFT01_DuplicateOp_Fires(t *testing.T) {
	fs := &yongol.Fullstack{
		Features: []featparser.Feature{
			{Op: "CreateWorkflow", Path: "POST /workflows", Desc: "Create", Line: 2},
			{Op: "CreateWorkflow", Path: "POST /workflows/v2", Desc: "Create v2", Line: 5},
		},
	}
	diags := ft01DuplicateOp(fs)
	if len(diags) != 1 {
		t.Fatalf("want 1 diag, got %d", len(diags))
	}
	if !strings.Contains(diags[0].Message, "[FT-01]") {
		t.Errorf("want [FT-01] prefix, got %s", diags[0].Message)
	}
}

func TestFT01_DuplicateOp_NoFire(t *testing.T) {
	fs := &yongol.Fullstack{
		Features: []featparser.Feature{
			{Op: "CreateWorkflow", Path: "POST /workflows", Desc: "Create", Line: 2},
			{Op: "GetWorkflow", Path: "GET /workflows/{id}", Desc: "Get", Line: 5},
		},
	}
	diags := ft01DuplicateOp(fs)
	if len(diags) != 0 {
		t.Fatalf("want 0 diags, got %d", len(diags))
	}
}
