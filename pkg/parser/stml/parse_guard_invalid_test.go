//ff:func feature=stml-parse type=test control=iteration dimension=1
//ff:what ParseGuard 문법 위반 입력이 오류를 반환하는지 검증 (TM-17 근거)

package stml

import "testing"

func TestParseGuardInvalid(t *testing.T) {
	cases := []string{
		"",                          // empty
		"foo()",                     // function call
		"workflow.status",           // ref without op/lifecycle
		"workflow.status =",         // missing value
		"workflow.status == draft",  // == leaves trailing token
		"(workflow.status=a",        // unclosed group
		"workflow.status=a &&",      // dangling combinator
		"status=draft",              // ref without model.field dot
		".loading",                  // leading dot (legacy form, not a guard)
		"workflow.status=a + b.y=2", // arithmetic
		"workflow.status.bogus",     // invalid lifecycle keyword
		"workflow.status=a # b",     // unexpected character
	}
	for _, c := range cases {
		if _, err := ParseGuard(c); err == nil {
			t.Errorf("ParseGuard(%q) expected error, got nil", c)
		}
	}
}
