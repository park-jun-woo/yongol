//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what tm30OutsideEachDiag — ERROR 레벨·파일·op·메시지 구성 검증

package stml_openapi

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

func TestTM30OutsideEachDiag(t *testing.T) {
	d := tm30OutsideEachDiag("p.html", "DeletePhoto", "item.id")
	if d.Level != diagnostic.LevelError || d.Phase != diagnostic.PhaseValidate {
		t.Errorf("level/phase: %+v", d)
	}
	if d.File != "p.html" || d.OperationID != "DeletePhoto" {
		t.Errorf("file/op: %+v", d)
	}
	if !strings.Contains(d.Message, "[TM-30]") || !strings.Contains(d.Message, `"item.id"`) {
		t.Errorf("message: %q", d.Message)
	}
	if d.Advice == "" {
		t.Error("advice must not be empty")
	}
}
