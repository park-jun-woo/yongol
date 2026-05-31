//ff:func feature=migration type=test control=sequence
//ff:what render_helpers_unit_test — appendColumnLines/PK/FK/Check/Index/Sentinel + renderTable + CanonicalSQL 단위 테스트
package migration

import (
	"strings"
	"testing"
)

func TestAppendPrimaryKeyClause(t *testing.T) {
	var b strings.Builder
	appendPrimaryKeyClause(&b, []string{"id", "tenant"})
	if got := b.String(); got != ",\n    PRIMARY KEY (id, tenant)" {
		t.Errorf("got %q", got)
	}

	var empty strings.Builder
	appendPrimaryKeyClause(&empty, nil)
	if empty.String() != "" {
		t.Errorf("empty pk should be no-op, got %q", empty.String())
	}
}
