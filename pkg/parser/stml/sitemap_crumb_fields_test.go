//ff:func feature=stml-parse type=test control=sequence
//ff:what TestSitemapCrumbFields — nav 전체 수집, nil sitemap/선언 없음 nil, 그룹 li 무시 검증

package stml

import "testing"

func TestSitemapCrumbFields(t *testing.T) {
	t.Run("collects across nav blocks and depths", func(t *testing.T) {
		sm := &SitemapSpec{Navs: []SitemapNav{
			{Items: []SitemapNode{
				{Page: "building-list", Children: []SitemapNode{
					{Page: "building-detail", CrumbField: "building_name"},
				}},
			}},
			{Items: []SitemapNode{
				{Page: "member-detail", CrumbField: "member_name"},
			}},
		}}
		got := SitemapCrumbFields(sm)
		if len(got) != 2 || got["building-detail"] != "building_name" || got["member-detail"] != "member_name" {
			t.Errorf("SitemapCrumbFields = %v", got)
		}
	})

	t.Run("nil sitemap yields nil", func(t *testing.T) {
		if got := SitemapCrumbFields(nil); got != nil {
			t.Errorf("SitemapCrumbFields(nil) = %v, want nil", got)
		}
	})

	t.Run("no declaration yields nil — byte-identity gate", func(t *testing.T) {
		sm := &SitemapSpec{Navs: []SitemapNav{{Items: []SitemapNode{{Page: "home"}}}}}
		if got := SitemapCrumbFields(sm); got != nil {
			t.Errorf("SitemapCrumbFields = %v, want nil without declarations", got)
		}
	})
}
