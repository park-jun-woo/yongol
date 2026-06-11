//ff:func feature=gen-react type=test control=sequence
//ff:what collectMenuRoleSets — 문서 순서 수집/상수명 중복 제거/role 없는 항목·빈 트리 검증

package react

import (
	"reflect"
	"testing"
)

func TestCollectMenuRoleSets(t *testing.T) {
	items := []sitemapMenuItem{
		{Kind: "page", Roles: []string{"admin", "manager"}},
		{Kind: "group", Roles: []string{"admin"}, Children: []sitemapMenuItem{
			{Kind: "page", Roles: []string{"admin"}}, // duplicate set, deduped
			{Kind: "page", Roles: []string{"member"}},
		}},
		{Kind: "page"}, // no roles
	}
	got := collectMenuRoleSets(items)
	want := [][]string{{"admin", "manager"}, {"admin"}, {"member"}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
	if sets := collectMenuRoleSets(nil); len(sets) != 0 {
		t.Errorf("nil items: got %v", sets)
	}
}
