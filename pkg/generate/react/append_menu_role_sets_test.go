//ff:func feature=gen-react type=test control=iteration dimension=1
//ff:what appendMenuRoleSets — 깊이 우선 순서/상수명 기준 중복 제거/seen 선등록 스킵 검증

package react

import (
	"reflect"
	"testing"
)

func TestAppendMenuRoleSets(t *testing.T) {
	items := []sitemapMenuItem{
		{Kind: "group", Roles: []string{"admin"}, Children: []sitemapMenuItem{
			{Kind: "page", Roles: []string{"admin", "manager"}}, // child visited after parent (depth-first)
			{Kind: "page", Roles: []string{"admin"}},            // duplicate constant name → skipped
		}},
		{Kind: "page"}, // no roles → not recorded
		{Kind: "page", Roles: []string{"member"}},
	}
	seen := map[string]bool{}
	var sets [][]string
	appendMenuRoleSets(items, seen, &sets)

	want := [][]string{{"admin"}, {"admin", "manager"}, {"member"}}
	if !reflect.DeepEqual(sets, want) {
		t.Errorf("sets = %v, want %v", sets, want)
	}
	for _, set := range want {
		if !seen[rolesConstName(set)] {
			t.Errorf("seen[%q] = false, want true", rolesConstName(set))
		}
	}

	// a constant name already in seen is not appended again
	var sets2 [][]string
	appendMenuRoleSets([]sitemapMenuItem{{Kind: "page", Roles: []string{"member"}}},
		map[string]bool{rolesConstName([]string{"member"}): true}, &sets2)
	if len(sets2) != 0 {
		t.Errorf("pre-seen set appended: %v", sets2)
	}
}
