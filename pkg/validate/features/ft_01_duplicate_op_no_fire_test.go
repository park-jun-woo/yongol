//ff:func feature=validate type=test control=sequence topic=features-structural
//ff:what FT-01 — op 중복 없을 때 정상 통과 테스트
package features

import (
	"testing"

	featparser "github.com/park-jun-woo/yongol/pkg/parser/features"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

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
