//ff:func feature=validate type=test control=sequence topic=openapi-ssac
//ff:what TestXos70_SubscribeFuncSkipped — subscribe 함수는 XOS-70 검증 스킵 확인

package openapi_ssac

import (
	"testing"

	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXos70_SubscribeFuncSkipped(t *testing.T) {
	fs := &yongol.Fullstack{
		ServiceFuncs: []ssac.ServiceFunc{
			{Name: "onOrder", Subscribe: &ssac.SubscribeInfo{}},
		},
		ResponseConstraints: map[string]map[string]oapiparser.FieldConstraint{},
	}
	diags := xos70ResponseLiteralIntFormat(fs)
	if len(diags) != 0 {
		t.Fatalf("expected 0, got %d", len(diags))
	}
}
