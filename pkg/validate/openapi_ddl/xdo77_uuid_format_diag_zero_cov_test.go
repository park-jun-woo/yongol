//ff:func feature=validate type=test control=sequence
//ff:what TestXdo77Diag_ZeroCov — float/uuid format 누락 진단 빌더 직접 호출
package openapi_ddl

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

func TestXdo77UUIDFormatDiag_ZeroCov(t *testing.T) {
	d := xdo77UUIDFormatDiag("users", "id", 7)
	if d.Level != diagnostic.LevelError || d.Line != 7 {
		t.Errorf("uuid diag = %+v", d)
	}
	if !strings.Contains(d.Message, "users.id") || !strings.Contains(d.Message, "format: uuid") {
		t.Errorf("uuid diag message = %q", d.Message)
	}
}
