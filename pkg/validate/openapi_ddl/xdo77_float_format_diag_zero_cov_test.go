//ff:func feature=validate type=test control=sequence
//ff:what TestXdo77Diag_ZeroCov — float/uuid format 누락 진단 빌더 직접 호출
package openapi_ddl

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

func TestXdo77FloatFormatDiag_ZeroCov(t *testing.T) {
	d := xdo77FloatFormatDiag("orders", "amount", 42)
	if d.Level != diagnostic.LevelError || d.Line != 42 {
		t.Errorf("float diag = %+v", d)
	}
	if !strings.Contains(d.Message, "orders.amount") || !strings.Contains(d.Message, "format: double") {
		t.Errorf("float diag message = %q", d.Message)
	}
}
