//ff:func feature=migration type=test control=sequence
//ff:what render_helpers_unit_test — appendColumnLines/PK/FK/Check/Index/Sentinel + renderTable + CanonicalSQL 단위 테스트
package migration

import (
	"strings"
	"testing"
)

func TestAppendSentinelLines(t *testing.T) {
	var b strings.Builder
	appendSentinelLines(&b, []SentinelInsert{
		{SQL: "INSERT INTO t VALUES (1);\n\n"},
		{SQL: "INSERT INTO t VALUES (2);"},
	})
	want := "\nINSERT INTO t VALUES (1);\n\nINSERT INTO t VALUES (2);\n"
	if got := b.String(); got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
