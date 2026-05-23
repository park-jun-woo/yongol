//ff:func feature=validate type=test control=sequence topic=hurl-openapi
//ff:what flushCurJSONPath — cur 축적된 식별자 flush/reset 검증

package hurl_openapi

import (
	"strings"
	"testing"
)

func TestFlushCurJSONPath(t *testing.T) {
	t.Run("empty_cur_no_append", func(t *testing.T) {
		var cur strings.Builder
		var out []string
		flushCurJSONPath(&cur, &out)
		if len(out) != 0 {
			t.Errorf("expected empty out, got %v", out)
		}
	})

	t.Run("non_empty_cur_appends_and_resets", func(t *testing.T) {
		var cur strings.Builder
		cur.WriteString("name")
		var out []string
		flushCurJSONPath(&cur, &out)
		if len(out) != 1 || out[0] != "name" {
			t.Errorf("out = %v, want [name]", out)
		}
		if cur.Len() != 0 {
			t.Errorf("cur should be reset, but has len %d", cur.Len())
		}
	})

	t.Run("multiple_flushes_accumulate", func(t *testing.T) {
		var cur strings.Builder
		var out []string
		cur.WriteString("a")
		flushCurJSONPath(&cur, &out)
		cur.WriteString("b")
		flushCurJSONPath(&cur, &out)
		if len(out) != 2 || out[0] != "a" || out[1] != "b" {
			t.Errorf("out = %v, want [a b]", out)
		}
	})
}
