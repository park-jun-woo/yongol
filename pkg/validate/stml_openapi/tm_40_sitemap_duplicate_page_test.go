//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what TM-40 — 블록 간 중복 등장 ERROR(두 위치 표기) / 단일 등장 침묵 검증

package stml_openapi

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestTM40SitemapDuplicatePage(t *testing.T) {
	pages := []stml.PageSpec{{Name: "login", FileName: "login.html"}}

	t.Run("duplicate across nav blocks fires with both positions", func(t *testing.T) {
		fs := makeFS(pages, nil)
		fs.Sitemap = &stml.SitemapSpec{
			FileName: "sitemap.html",
			Navs: []stml.SitemapNav{
				{Items: []stml.SitemapNode{{Page: "login", Label: "로그인", Menu: true}}},
				{Items: []stml.SitemapNode{{Label: "인증", Menu: true, Children: []stml.SitemapNode{
					{Page: "login", Label: "로그인 다시", Menu: true},
				}}}},
			},
		}
		diags := tm40SitemapDuplicatePage(fs)
		if got := countDiag(diags, "[TM-40]"); got != 1 {
			t.Fatalf("expected 1 TM-40, got %d: %+v", got, diags)
		}
		if diags[0].Level != diagnostic.LevelError {
			t.Errorf("Level = %v, want LevelError", diags[0].Level)
		}
		if !strings.Contains(diags[0].Message, "nav[0] > 로그인") || !strings.Contains(diags[0].Message, "nav[1] > 인증 > 로그인 다시") {
			t.Errorf("Message should show both positions, got %q", diags[0].Message)
		}
	})

	t.Run("single appearance is silent", func(t *testing.T) {
		fs := makeFS(pages, nil)
		fs.Sitemap = &stml.SitemapSpec{
			FileName: "sitemap.html",
			Navs: []stml.SitemapNav{
				{Items: []stml.SitemapNode{{Page: "login", Label: "로그인", Menu: true}}},
			},
		}
		if diags := tm40SitemapDuplicatePage(fs); len(diags) != 0 {
			t.Errorf("expected silence, got %+v", diags)
		}
	})
}
