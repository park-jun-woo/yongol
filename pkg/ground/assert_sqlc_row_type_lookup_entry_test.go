//ff:func feature=rule type=test-helper control=sequence
//ff:what SQLc.rowType.<want> Lookup 에 want 가 등록되었는지 assert

package ground

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/rule"
)

// assertSQLcRowTypeLookupEntry reports a test error when
// g.Lookup["SQLc.rowType."+want] is missing or does not contain want.
func assertSQLcRowTypeLookupEntry(t *testing.T, g *rule.Ground, want string) {
	t.Helper()
	set, ok := g.Lookup["SQLc.rowType."+want]
	if !ok {
		t.Errorf("SQLc.rowType.%s missing; Lookup=%v", want, g.Lookup)
		return
	}
	if !set[want] {
		t.Errorf("SQLc.rowType.%s set does not contain %q: %v", want, want, set)
	}
}
