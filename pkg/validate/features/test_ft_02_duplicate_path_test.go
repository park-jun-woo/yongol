//ff:what FT-02 — path 중복 시 에러 진단 테스트
package features

import (
	"strings"
	"testing"

	featparser "github.com/park-jun-woo/yongol/pkg/parser/features"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestFT02_DuplicatePath_Fires(t *testing.T) {
	fs := &yongol.Fullstack{
		Features: []featparser.Feature{
			{Op: "CreateWorkflow", Path: "POST /workflows", Desc: "Create", Line: 2},
			{Op: "DuplicateWorkflow", Path: "POST /workflows", Desc: "Duplicate", Line: 5},
		},
	}
	diags := ft02DuplicatePath(fs)
	if len(diags) != 1 {
		t.Fatalf("want 1 diag, got %d", len(diags))
	}
	if !strings.Contains(diags[0].Message, "[FT-02]") {
		t.Errorf("want [FT-02] prefix, got %s", diags[0].Message)
	}
}

func TestFT02_DuplicatePath_NoFire(t *testing.T) {
	fs := &yongol.Fullstack{
		Features: []featparser.Feature{
			{Op: "CreateWorkflow", Path: "POST /workflows", Desc: "Create", Line: 2},
			{Op: "GetWorkflow", Path: "GET /workflows/{id}", Desc: "Get", Line: 5},
		},
	}
	diags := ft02DuplicatePath(fs)
	if len(diags) != 0 {
		t.Fatalf("want 0 diags, got %d", len(diags))
	}
}
