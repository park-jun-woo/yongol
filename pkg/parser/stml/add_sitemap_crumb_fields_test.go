//ff:func feature=stml-parse type=test control=sequence
//ff:what TestAddSitemapCrumbFields — 지연 할당/그룹 li 비기록/최초 등장 우선 검증

package stml

import "testing"

func TestAddSitemapCrumbFields(t *testing.T) {
	t.Run("lazy allocation stays nil without declarations", func(t *testing.T) {
		got := addSitemapCrumbFields([]SitemapNode{{Page: "home"}, {Label: "그룹"}}, nil)
		if got != nil {
			t.Errorf("expected nil map, got %v", got)
		}
	})

	t.Run("page-less node records nothing even with the attribute", func(t *testing.T) {
		got := addSitemapCrumbFields([]SitemapNode{{Label: "그룹", CrumbField: "name"}}, nil)
		if got != nil {
			t.Errorf("expected nil map for a group li, got %v", got)
		}
	})

	t.Run("first occurrence wins", func(t *testing.T) {
		got := addSitemapCrumbFields([]SitemapNode{
			{Page: "detail", CrumbField: "first"},
			{Page: "detail", CrumbField: "second"},
		}, nil)
		if len(got) != 1 || got["detail"] != "first" {
			t.Errorf("got %v, want detail → first", got)
		}
	})

	t.Run("recurses into children", func(t *testing.T) {
		got := addSitemapCrumbFields([]SitemapNode{
			{Label: "그룹", Children: []SitemapNode{{Page: "detail", CrumbField: "name"}}},
		}, nil)
		if len(got) != 1 || got["detail"] != "name" {
			t.Errorf("got %v, want detail → name", got)
		}
	})
}
