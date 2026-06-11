//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what tm53Diag — TM-53 WARNING 진단 구성(파일·레벨·메시지·advice) 검증

package stml_openapi

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

func TestTM53Diag(t *testing.T) {
	d := tm53Diag("p.html", "[TM-53] msg", "do this")
	if len(d) != 1 {
		t.Fatalf("expected 1 diagnostic, got %d", len(d))
	}
	if d[0].File != "p.html" || d[0].Level != diagnostic.LevelWarning ||
		d[0].Phase != diagnostic.PhaseValidate || d[0].Message != "[TM-53] msg" || d[0].Advice != "do this" {
		t.Errorf("unexpected diagnostic: %+v", d[0])
	}
}
