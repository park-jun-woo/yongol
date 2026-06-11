//ff:func feature=gen-react type=test control=sequence
//ff:what menuIconNames — 중첩 수집/정렬/중복 제거/빈 입력 nil 검증

package react

import (
	"reflect"
	"testing"
)

func TestMenuIconNames(t *testing.T) {
	items := []sitemapMenuItem{
		{Kind: "page", Icon: "LayoutDashboard"},
		{Kind: "group", Icon: "Building", Children: []sitemapMenuItem{
			{Kind: "page", Icon: "Building"}, // duplicate
			{Kind: "page", Icon: "FileText"},
		}},
	}
	want := []string{"Building", "FileText", "LayoutDashboard"}
	if got := menuIconNames(items); !reflect.DeepEqual(got, want) {
		t.Errorf("menuIconNames = %v, want %v", got, want)
	}

	if got := menuIconNames([]sitemapMenuItem{{Kind: "page"}}); got != nil {
		t.Errorf("no icons must be nil, got %v", got)
	}
}
