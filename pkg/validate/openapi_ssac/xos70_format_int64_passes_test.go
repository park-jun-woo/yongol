//ff:func feature=validate type=test control=sequence topic=openapi-ssac
//ff:what TestXos70_FormatInt64Passes — format: int64 지정 시 진단 미발생 확인

package openapi_ssac

import (
	"testing"

	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXos70_FormatInt64Passes(t *testing.T) {
	fs := &yongol.Fullstack{
		ServiceFuncs: []ssac.ServiceFunc{
			{Name: "getUser", Sequences: []ssac.Sequence{
				{Type: "response", Fields: map[string]string{"count": "0"}},
			}},
		},
		ResponseConstraints: map[string]map[string]oapiparser.FieldConstraint{
			"getUser": {"count": {Type: "integer", Format: "int64", Required: false}},
		},
	}
	diags := xos70ResponseLiteralIntFormat(fs)
	if len(diags) != 0 {
		t.Fatalf("expected 0, got %d", len(diags))
	}
}
