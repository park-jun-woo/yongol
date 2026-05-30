//ff:func feature=contract type=test control=iteration dimension=1
//ff:what test: TestNormalizeBody — CRLF/BOM/trailing newline 정규화 테이블 주도 검증

package contract

import (
	"bytes"
	"testing"
)

func TestNormalizeBody(t *testing.T) {
	cases := []struct {
		name string
		in   []byte
		want []byte
	}{
		{"empty", []byte{}, []byte{}},
		{"crlf", []byte("a\r\nb\r\n"), []byte("a\nb\n")},
		{"bare_cr", []byte("a\rb\r"), []byte("a\nb\n")},
		{"bom_stripped", append([]byte{0xEF, 0xBB, 0xBF}, []byte("pkg\n")...), []byte("pkg\n")},
		{"append_trailing_lf", []byte("pkg"), []byte("pkg\n")},
		{"already_lf", []byte("a\nb\n"), []byte("a\nb\n")},
		{"only_bom", []byte{0xEF, 0xBB, 0xBF}, []byte{}},
	}
	for _, c := range cases {
		got := NormalizeBody(c.in)
		if !bytes.Equal(got, c.want) {
			t.Errorf("%s: got %q want %q", c.name, got, c.want)
		}
	}
}
