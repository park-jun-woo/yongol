//ff:func feature=gen-react type=test control=sequence
//ff:what renderSitemapItemBody — leaf 한 줄 <li>/자식 중첩 <ul> 들여쓰기/role-gated 자식의 조건 래핑 검증

package react

import (
	"strings"
	"testing"
)

func TestRenderSitemapItemBody(t *testing.T) {
	// leaf item → a single <li> line at the given indent
	var leaf strings.Builder
	renderSitemapItemBody(&leaf, sitemapMenuItem{Kind: "page", Label: "홈", To: "/"}, "    ", false)
	if got, want := leaf.String(), "    <li><NavLink to=\"/\" end>홈</NavLink></li>\n"; got != want {
		t.Errorf("leaf:\ngot  %q\nwant %q", got, want)
	}

	// group with children → expanded <li> with a nested <ul>, children
	// indented four more spaces (ul + li)
	var grp strings.Builder
	renderSitemapItemBody(&grp, sitemapMenuItem{Kind: "group", Label: "관리", Children: []sitemapMenuItem{
		{Kind: "page", Label: "회원", To: "/members"},
	}}, "  ", false)
	want := "  <li>\n" +
		"    <span>관리</span>\n" +
		"    <ul>\n" +
		"      <li><NavLink to=\"/members\" end>회원</NavLink></li>\n" +
		"    </ul>\n" +
		"  </li>\n"
	if got := grp.String(); got != want {
		t.Errorf("group:\ngot  %q\nwant %q", got, want)
	}

	// a role-gated child recurses through renderSitemapItem, so it wraps
	// itself in the conditional render when roles are active
	var gated strings.Builder
	renderSitemapItemBody(&gated, sitemapMenuItem{Kind: "group", Label: "관리", Children: []sitemapMenuItem{
		{Kind: "page", Label: "정산", To: "/billing", Roles: []string{"admin"}},
	}}, "", true)
	assertContains(t, gated.String(), "    {ROLES_admin.includes(userRole) && (\n")
	assertContains(t, gated.String(), "      <li><NavLink to=\"/billing\" end>정산</NavLink></li>\n")
}
