//ff:func feature=external type=test control=iteration dimension=1
//ff:what TestWriteBodyMap/writeReturnWithResult/readBodySnippet — 코드생성·스니펫 검증
package external

import (
	"strings"
	"testing"
)

func TestReadBodySnippet(t *testing.T) {
	tests := []struct {
		name  string
		in    string
		limit int
		want  string
	}{
		{"short", "hello", 100, "hello"},
		{"trimmed", "  hi  ", 100, "hi"},
		{"newlines collapsed", "a\nb\nc", 100, "a b c"},
		{"truncated", "abcdefgh", 3, "abc..."},
		{"empty", "", 10, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := readBodySnippet(strings.NewReader(tt.in), tt.limit)
			if got != tt.want {
				t.Errorf("readBodySnippet(%q, %d) = %q, want %q", tt.in, tt.limit, got, tt.want)
			}
		})
	}
}
