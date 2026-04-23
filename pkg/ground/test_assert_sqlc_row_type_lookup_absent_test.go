//ff:func feature=rule type=test-helper control=sequence
//ff:what exec/execresult 쿼리는 SQLc.rowType.* 에 미등록임을 assert

package ground

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/rule"
)

// assertSQLcRowTypeLookupAbsent reports a test error when an
// exec/execresult query is registered under SQLc.rowType.<noRow[Row]>.
func assertSQLcRowTypeLookupAbsent(t *testing.T, g *rule.Ground, noRow string) {
	t.Helper()
	if _, ok := g.Lookup["SQLc.rowType."+noRow]; ok {
		t.Errorf("SQLc.rowType.%s unexpectedly registered for exec/execresult", noRow)
	}
	if _, ok := g.Lookup["SQLc.rowType."+noRow+"Row"]; ok {
		t.Errorf("SQLc.rowType.%sRow unexpectedly registered for exec/execresult", noRow)
	}
}
