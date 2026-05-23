//ff:func feature=validate type=test control=sequence topic=hurl-openapi
//ff:what parseJSONPathChar — 개별 문자 처리 (dot/bracket/regular) 검증

package hurl_openapi

import (
	"strings"
	"testing"
)

func TestParseJSONPathChar(t *testing.T) {
	t.Run("dot_flushes_cur", func(t *testing.T) {
		var cur strings.Builder
		cur.WriteString("name")
		var out []string
		ret := parseJSONPathChar(".", 0, &cur, &out)
		if ret != 0 {
			t.Errorf("return = %d, want 0", ret)
		}
		if len(out) != 1 || out[0] != "name" {
			t.Errorf("out = %v, want [name]", out)
		}
		if cur.Len() != 0 {
			t.Errorf("cur should be reset")
		}
	})

	t.Run("bracket_consumes_index", func(t *testing.T) {
		var cur strings.Builder
		var out []string
		p := "[0].next"
		ret := parseJSONPathChar(p, 0, &cur, &out)
		if ret != 2 { // index of ']' in "[0].next" is 2
			t.Errorf("return = %d, want 2", ret)
		}
		if len(out) != 1 || out[0] != "[0]" {
			t.Errorf("out = %v, want [[0]]", out)
		}
	})

	t.Run("bracket_no_close_returns_end", func(t *testing.T) {
		var cur strings.Builder
		var out []string
		p := "[broken"
		ret := parseJSONPathChar(p, 0, &cur, &out)
		if ret != len(p)-1 {
			t.Errorf("return = %d, want %d", ret, len(p)-1)
		}
	})

	t.Run("regular_char_appended", func(t *testing.T) {
		var cur strings.Builder
		var out []string
		ret := parseJSONPathChar("abc", 0, &cur, &out)
		if ret != 0 {
			t.Errorf("return = %d, want 0", ret)
		}
		if cur.String() != "a" {
			t.Errorf("cur = %q, want \"a\"", cur.String())
		}
		if len(out) != 0 {
			t.Errorf("out should be empty")
		}
	})

	t.Run("bracket_flushes_pending_cur", func(t *testing.T) {
		var cur strings.Builder
		cur.WriteString("items")
		var out []string
		p := "[0]"
		parseJSONPathChar(p, 0, &cur, &out)
		if len(out) != 2 || out[0] != "items" || out[1] != "[0]" {
			t.Errorf("out = %v, want [items [0]]", out)
		}
	})
}
