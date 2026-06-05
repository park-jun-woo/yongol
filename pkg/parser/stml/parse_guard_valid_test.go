//ff:func feature=stml-parse type=test control=iteration dimension=1
//ff:what ParseGuard 정상 입력이 오류 없이 파싱되는지 검증

package stml

import "testing"

func TestParseGuardValid(t *testing.T) {
	cases := []string{
		"workflow.status=draft",
		"workflow.status != active",
		"order.count >= 3",
		"order.count < 10",
		"workflow.status=active && currentUser.Role=owner",
		"a.x=1 || b.y=2",
		"!workflow.locked=true",
		"(workflow.status=active || workflow.status=draft) && currentUser.Role=owner",
		"items.list.empty",
		"workflow.status = 'multi word'",
	}
	for _, c := range cases {
		if _, err := ParseGuard(c); err != nil {
			t.Errorf("ParseGuard(%q) unexpected error: %v", c, err)
		}
	}
}
