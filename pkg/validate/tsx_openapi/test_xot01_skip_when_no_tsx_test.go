//ff:func feature=validate type=test control=sequence topic=tsx-openapi
//ff:what XOT-01 skip — TSX 페이지가 없으면 규칙 비활성

package tsx_openapi

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/rule"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXot01_SkipWhenNoTSX(t *testing.T) {
	fs := &yongol.Fullstack{}
	fs.SetGround(&rule.Ground{Lookup: map[string]rule.StringSet{
		"OpenAPI.operationId": {"any": true},
	}})
	if diags := xot01OperationID(fs); len(diags) != 0 {
		t.Errorf("expected no diagnostics when no TSX pages, got %+v", diags)
	}
}
