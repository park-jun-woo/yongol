//ff:func feature=manifest type=test control=sequence
//ff:what extractParenColumns — UNIQUE (a, b, c) 괄호 내 컬럼 3개 추출

package ddl

import "testing"

func TestExtractParenColumns_Multi(t *testing.T) {
	got := extractParenColumns("UNIQUE (a, b , c)")
	if len(got) != 3 || got[0] != "a" || got[1] != "b" || got[2] != "c" {
		t.Errorf("got %v", got)
	}
}
