//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what TM-42 — 필수 세그먼트 라우트(TM-34 판정) / manifest.frontend.index 모순 발화와 일치·TM-39 영역 침묵 검증

package stml_openapi

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestTM42SitemapIndexConflict_Cases(t *testing.T) {
	pages := []stml.PageSpec{
		{Name: "dashboard", FileName: "dashboard.html"},
		{Name: "login", FileName: "login.html"},
	}

	t.Run("required segment route fires (TM-34 judgment)", func(t *testing.T) {
		detail := stml.PageSpec{
			Name:     "building-detail",
			FileName: "building-detail.html",
			Fetches: []stml.FetchBlock{{
				OperationID: "GetBuilding",
				Params:      []stml.ParamBind{{Name: "id", Source: "route.BuildingID"}},
			}},
		}
		fs := makeFS(append(pages, detail), nil)
		fs.Sitemap = sitemapWithIndexFixture("building-detail")
		diags := tm42SitemapIndexConflict(fs)
		if got := countDiag(diags, "[TM-42]"); got != 1 {
			t.Fatalf("expected 1 TM-42 for required segment, got %d: %+v", got, diags)
		}
		if !strings.Contains(diags[0].Message, ":BuildingID") {
			t.Errorf("Message should name the required segment, got %q", diags[0].Message)
		}
	})

	t.Run("manifest.frontend.index contradiction fires", func(t *testing.T) {
		fs := makeFS(pages, nil)
		fs.Manifest.Frontend.Index = "login"
		fs.Sitemap = sitemapWithIndexFixture("dashboard")
		diags := tm42SitemapIndexConflict(fs)
		if got := countDiag(diags, "[TM-42]"); got != 1 {
			t.Fatalf("expected 1 TM-42 for the contradiction, got %d: %+v", got, diags)
		}
		if !strings.Contains(diags[0].Message, `"dashboard"`) || !strings.Contains(diags[0].Message, `"login"`) {
			t.Errorf("Message should name both declarations, got %q", diags[0].Message)
		}
	})

	t.Run("manifest.frontend.index agreement is silent", func(t *testing.T) {
		fs := makeFS(pages, nil)
		fs.Manifest.Frontend.Index = "dashboard"
		fs.Sitemap = sitemapWithIndexFixture("dashboard")
		if diags := tm42SitemapIndexConflict(fs); len(diags) != 0 {
			t.Errorf("expected silence when both name the same page, got %+v", diags)
		}
	})

	t.Run("nonexistent index page is TM-39 territory", func(t *testing.T) {
		fs := makeFS(pages, nil)
		fs.Sitemap = sitemapWithIndexFixture("ghost")
		if diags := tm42SitemapIndexConflict(fs); len(diags) != 0 {
			t.Errorf("expected silence (TM-39 reports the missing page), got %+v", diags)
		}
	})
}
