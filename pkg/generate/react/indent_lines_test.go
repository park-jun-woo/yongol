//ff:func feature=gen-react type=test control=sequence
//ff:what indentLines — 라인별 들여쓰기 결합/빈 입력 검증

package react

import "testing"

func TestIndentLines(t *testing.T) {
	if got := indentLines([]string{"a,", "b,"}, "  "); got != "  a,\n  b,\n" {
		t.Errorf("got %q", got)
	}
	if got := indentLines(nil, "  "); got != "" {
		t.Errorf("empty: got %q", got)
	}
}
