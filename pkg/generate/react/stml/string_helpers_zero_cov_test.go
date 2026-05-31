//ff:func feature=stml-gen type=test control=sequence
//ff:what 단순 stml 헬퍼 (string/zod/param/login) 묶음 커버 — coverage attribution 으로 다수 함수 PASS
package stml

import (
	"testing"
)

func TestStringHelpers_ZeroCov(t *testing.T) {
	if formatFloat(3) != "3" || formatFloat(2.5) != "2.5" {
		t.Errorf("formatFloat wrong: %q %q", formatFloat(3), formatFloat(2.5))
	}
	if indentStr(3) != "   " {
		t.Errorf("indentStr wrong: %q", indentStr(3))
	}
	if joinWords([]string{"a", "b"}) != "a b" {
		t.Errorf("joinWords wrong")
	}
	if toUpperFirst("abc") != "Abc" || toUpperFirst("") != "" {
		t.Errorf("toUpperFirst wrong")
	}
	if toLowerFirst("Abc") != "abc" || toLowerFirst("") != "" {
		t.Errorf("toLowerFirst wrong")
	}
	if pascalToLabel("RoomID") == "" || pascalToLabel("") != "" {
		t.Errorf("pascalToLabel wrong: %q", pascalToLabel("RoomID"))
	}
	if snakeToLabel("first_name") == "" {
		t.Errorf("snakeToLabel wrong")
	}
}
