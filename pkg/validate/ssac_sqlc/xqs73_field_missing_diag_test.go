//ff:func feature=validate type=test control=sequence topic=ssac-sqlc
//ff:what TestSSaCSqlcHelpers — unit tests for the pure ssac_sqlc helper functions
package ssac_sqlc

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	sqlcparser "github.com/park-jun-woo/yongol/pkg/parser/sqlc"
	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestXqs73FieldMissingDiag(t *testing.T) {
	fn := ssacparser.ServiceFunc{FileName: "svc.ssac"}
	seq := ssacparser.Sequence{Line: 42}
	vi := xqs73VarInfo{query: sqlcparser.QuerySpec{Name: "GetUser", SelectCols: []string{"id", "email"}}}
	d := xqs73FieldMissingDiag(fn, seq, "userName", "u", vi)
	if d.File != "svc.ssac" || d.Line != 42 {
		t.Errorf("loc = (%q,%d), want (svc.ssac,42)", d.File, d.Line)
	}
	if d.Level != diagnostic.LevelError || d.Phase != diagnostic.PhaseValidate {
		t.Errorf("level/phase mismatch: %v/%v", d.Level, d.Phase)
	}
	if !strings.Contains(d.Message, "[XQS-73]") || !strings.Contains(d.Message, "userName") ||
		!strings.Contains(d.Message, "GetUser") || !strings.Contains(d.Message, "id, email") {
		t.Errorf("message missing parts: %q", d.Message)
	}
	if !strings.Contains(d.Advice, "user_name") || !strings.Contains(d.Advice, "GetUser") {
		t.Errorf("advice missing parts: %q", d.Advice)
	}
}
