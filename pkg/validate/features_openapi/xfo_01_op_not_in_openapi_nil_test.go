//ff:func feature=validate type=test control=sequence topic=features-openapi
//ff:what XFO-01 — Ground nil 시 단락 테스트

package features_openapi

import (
	"testing"

	featparser "github.com/park-jun-woo/yongol/pkg/parser/features"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXFO01_OpNotInOpenAPI_NilGround(t *testing.T) {
	fs := &yongol.Fullstack{
		Features: []featparser.Feature{
			{Op: "CreateWorkflow", Path: "POST /workflows", Desc: "Create", Line: 2},
		},
	}
	diags := xfo01OpNotInOpenAPI(fs)
	if len(diags) != 0 {
		t.Fatalf("want 0 diags with nil ground, got %d", len(diags))
	}
}
