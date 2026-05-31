//ff:func feature=migration type=test control=sequence
//ff:what render_helpers_unit_test — appendColumnLines/PK/FK/Check/Index/Sentinel + renderTable + CanonicalSQL 단위 테스트
package migration

import (
	"strings"
	"testing"
)

func TestAppendCheckClauses(t *testing.T) {
	var b strings.Builder
	appendCheckClauses(&b, []*CheckConstraint{{Name: "chk_pos", Expression: "x > 0"}})
	if got := b.String(); got != ",\n    CONSTRAINT chk_pos CHECK (x > 0)" {
		t.Errorf("got %q", got)
	}
}
