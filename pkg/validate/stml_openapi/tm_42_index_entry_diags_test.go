//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what tm42IndexEntryDiags 단위 — page 부재/TM-39 영역 침묵/필수 세그먼트/manifest 모순(동시 발화 포함)/nil manifest 검증

package stml_openapi

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestTM42IndexEntryDiags(t *testing.T) {
	pages := []stml.PageSpec{
		{Name: "dashboard", FileName: "dashboard.html"},
		{Name: "building-detail", FileName: "building-detail.html",
			Fetches: []stml.FetchBlock{{
				OperationID: "GetBuilding",
				Params:      []stml.ParamBind{{Name: "id", Source: "route.BuildingID"}},
			}}},
	}
	entry := func(page, path string) sitemapEntry {
		return sitemapEntry{Node: stml.SitemapNode{Page: page, Label: page, Index: true, Menu: true}, Path: path}
	}

	t.Run("entry without data-page fires with its position", func(t *testing.T) {
		fs := makeFS(pages, nil)
		fs.Sitemap = sitemapWithIndexFixture("dashboard")
		diags := tm42IndexEntryDiags(fs, sitemapEntry{Node: stml.SitemapNode{Index: true, Label: "그룹"}, Path: "nav[0] > 그룹"})
		if len(diags) != 1 || diags[0].Level != diagnostic.LevelError {
			t.Fatalf("expected 1 error diag, got %+v", diags)
		}
		if !strings.Contains(diags[0].Message, "without data-page") || !strings.Contains(diags[0].Message, "nav[0] > 그룹") {
			t.Errorf("Message = %q, want the missing data-page named at its position", diags[0].Message)
		}
	})

	t.Run("nonexistent page stays silent (TM-39 territory)", func(t *testing.T) {
		fs := makeFS(pages, nil)
		fs.Sitemap = sitemapWithIndexFixture("ghost")
		if diags := tm42IndexEntryDiags(fs, entry("ghost", "nav[0] > ghost")); diags != nil {
			t.Errorf("expected nil, got %+v", diags)
		}
	})

	t.Run("required route segment fires (TM-34 judgment)", func(t *testing.T) {
		fs := makeFS(pages, nil)
		fs.Sitemap = sitemapWithIndexFixture("building-detail")
		diags := tm42IndexEntryDiags(fs, entry("building-detail", "nav[0] > building-detail"))
		if got := countDiag(diags, "[TM-42]"); got != 1 {
			t.Fatalf("expected 1 TM-42, got %d: %+v", got, diags)
		}
		if !strings.Contains(diags[0].Message, ":BuildingID") {
			t.Errorf("Message = %q, want the required segment named", diags[0].Message)
		}
	})

	t.Run("manifest.frontend.index contradiction fires", func(t *testing.T) {
		fs := makeFS(pages, nil)
		fs.Sitemap = sitemapWithIndexFixture("dashboard")
		fs.Manifest.Frontend.Index = "login"
		diags := tm42IndexEntryDiags(fs, entry("dashboard", "nav[0] > dashboard"))
		if got := countDiag(diags, "[TM-42]"); got != 1 {
			t.Fatalf("expected 1 TM-42, got %d: %+v", got, diags)
		}
		if !strings.Contains(diags[0].Message, `"dashboard"`) || !strings.Contains(diags[0].Message, `"login"`) {
			t.Errorf("Message = %q, want both declarations named", diags[0].Message)
		}
	})

	t.Run("required segment and manifest contradiction both fire", func(t *testing.T) {
		fs := makeFS(pages, nil)
		fs.Sitemap = sitemapWithIndexFixture("building-detail")
		fs.Manifest.Frontend.Index = "dashboard"
		diags := tm42IndexEntryDiags(fs, entry("building-detail", "nav[0] > building-detail"))
		if got := countDiag(diags, "[TM-42]"); got != 2 {
			t.Fatalf("expected 2 TM-42, got %d: %+v", got, diags)
		}
	})

	t.Run("valid entry is silent, nil manifest tolerated", func(t *testing.T) {
		fs := makeFS(pages, nil)
		fs.Sitemap = sitemapWithIndexFixture("dashboard")
		fs.Manifest = nil
		if diags := tm42IndexEntryDiags(fs, entry("dashboard", "nav[0] > dashboard")); len(diags) != 0 {
			t.Errorf("expected silence, got %+v", diags)
		}
	})

	t.Run("matching manifest index is silent", func(t *testing.T) {
		fs := makeFS(pages, nil)
		fs.Sitemap = sitemapWithIndexFixture("dashboard")
		fs.Manifest.Frontend.Index = "dashboard"
		if diags := tm42IndexEntryDiags(fs, entry("dashboard", "nav[0] > dashboard")); len(diags) != 0 {
			t.Errorf("expected silence, got %+v", diags)
		}
	})
}
