//ff:func feature=migration type=test control=sequence
//ff:what TestTrimForErr — 개행 치환 + 80자 초과 시 잘라내고 ... 부착
package migration

import (
	"strings"
	"testing"
)

func TestTrimForErr(t *testing.T) {
	if got := trimForErr("a\nb"); got != "a b" {
		t.Errorf("newline replace = %q, want %q", got, "a b")
	}
	short := "SELECT 1;"
	if got := trimForErr(short); got != short {
		t.Errorf("short unchanged = %q, want %q", got, short)
	}
	long := strings.Repeat("x", 100)
	got := trimForErr(long)
	if len(got) != 83 || !strings.HasSuffix(got, "...") {
		t.Errorf("long trim = %q (len %d), want 80 chars + ...", got, len(got))
	}
	if got[:80] != long[:80] {
		t.Errorf("prefix mismatch")
	}
}
