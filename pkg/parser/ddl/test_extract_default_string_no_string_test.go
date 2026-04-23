//ff:func feature=manifest type=test control=sequence
//ff:what extractDefaultString — 숫자 default 는 빈 문자열 반환

package ddl

import "testing"

func TestExtractDefaultString_NoString(t *testing.T) {
	if got := extractDefaultString("count INTEGER NOT NULL DEFAULT 0"); got != "" {
		t.Errorf("got %q, want empty (numeric default)", got)
	}
}
