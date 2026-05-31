//ff:func feature=agent type=test control=sequence
//ff:what TestIndentText — 비어있지 않은 줄에만 prefix 가 붙는지 검증
package agent

import (
	"testing"
)

func TestIndentText(t *testing.T) {
	got := indentText("a\n\nb", "  ")
	want := "  a\n\n  b\n"
	if got != want {
		t.Errorf("indentText = %q, want %q", got, want)
	}

	// Trailing newline is trimmed before processing.
	got = indentText("x\n", ">")
	if got != ">x\n" {
		t.Errorf("indentText trailing = %q, want %q", got, ">x\n")
	}
}
