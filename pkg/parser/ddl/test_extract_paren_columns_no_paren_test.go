//ff:func feature=manifest type=test control=sequence
//ff:what extractParenColumns — 괄호 없으면 nil 반환

package ddl

import "testing"

func TestExtractParenColumns_NoParen(t *testing.T) {
	if got := extractParenColumns("PRIMARY KEY id"); got != nil {
		t.Errorf("got %v, want nil", got)
	}
}
