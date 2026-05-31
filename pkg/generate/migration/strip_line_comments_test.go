//ff:func feature=migration type=test control=sequence
//ff:what TestStripLineComments — -- 라인 코멘트 제거, single-quote 내부 보존
package migration

import (
	"strings"
	"testing"
)

func TestStripLineComments(t *testing.T) {
	in := "SELECT 1; -- trailing\nSELECT '-- not a comment';\n"
	got := stripLineComments(in)
	if strings.Contains(got, "trailing") {
		t.Errorf("comment not stripped: %q", got)
	}
	if !strings.Contains(got, "'-- not a comment'") {
		t.Errorf("string literal damaged: %q", got)
	}
}
