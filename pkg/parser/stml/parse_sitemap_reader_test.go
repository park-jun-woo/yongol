//ff:func feature=stml-parse type=test control=sequence
//ff:what TestParseSitemapReader_FullGrammar — 그룹/외부 링크/중첩/라벨/data-* 속성/2개 nav 블록 파싱 검증

package stml

import (
	"strings"
	"testing"
)

func TestParseSitemapReader_FullGrammar(t *testing.T) {
	spec, diags := ParseSitemapReader("sitemap.html", strings.NewReader(sitemapFullGrammarSrc))
	if len(diags) != 0 {
		t.Fatalf("expected no diags, got %+v", diags)
	}
	if spec.FileName != "sitemap.html" {
		t.Errorf("FileName = %q, want sitemap.html", spec.FileName)
	}
	if len(spec.Navs) != 2 {
		t.Fatalf("expected 2 navs, got %d", len(spec.Navs))
	}

	app := spec.Navs[0]
	if app.Layout != "app" || app.Entry {
		t.Errorf("nav[0] = layout %q entry %v, want app/false", app.Layout, app.Entry)
	}
	if len(app.Items) != 4 {
		t.Fatalf("nav[0] expected 4 items, got %d: %+v", len(app.Items), app.Items)
	}

	dash := app.Items[0]
	if dash.Page != "dashboard" || dash.Label != "대시보드" || !dash.Index || !dash.Menu {
		t.Errorf("dashboard item = %+v", dash)
	}

	group := app.Items[1]
	if group.Page != "" || group.Href != "" || group.Label != "건물 관리" {
		t.Errorf("group item = %+v, want page-less label 건물 관리", group)
	}
	if len(group.Children) != 1 || group.Children[0].Page != "building-list" {
		t.Fatalf("group children = %+v", group.Children)
	}
	if kids := group.Children[0].Children; len(kids) != 1 || kids[0].Page != "building-detail" || kids[0].Label != "건물 상세" {
		t.Errorf("nested children = %+v", group.Children[0].Children)
	}

	member := app.Items[2]
	if member.Page != "member-list" || member.Icon != "users" || member.Menu {
		t.Errorf("member item = %+v, want icon users + menu false", member)
	}

	ext := app.Items[3]
	if ext.Page != "" || ext.Href != "https://docs.example.com" || ext.Label != "사용자 매뉴얼" {
		t.Errorf("external link item = %+v", ext)
	}

	bare := spec.Navs[1]
	if bare.Layout != "bare" || !bare.Entry {
		t.Errorf("nav[1] = layout %q entry %v, want bare/true", bare.Layout, bare.Entry)
	}
	if len(bare.Items) != 1 || bare.Items[0].Page != "login" {
		t.Errorf("nav[1] items = %+v", bare.Items)
	}
}
