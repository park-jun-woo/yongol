//ff:func feature=validate type=test control=sequence topic=openapi-ssac
//ff:what TestXos70_NilResponseConstraints — nil ResponseConstraints 시 빈 결과 검증

package openapi_ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXos70_NilResponseConstraints(t *testing.T) {
	fs := &yongol.Fullstack{
		ServiceFuncs: []ssac.ServiceFunc{
			{Name: "getUser", Sequences: []ssac.Sequence{
				{Type: "response", Fields: map[string]string{"count": "0"}},
			}},
		},
	}
	diags := xos70ResponseLiteralIntFormat(fs)
	if len(diags) != 0 {
		t.Fatalf("expected 0, got %d", len(diags))
	}
}
