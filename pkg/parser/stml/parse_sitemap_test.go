//ff:func feature=stml-parse type=test control=sequence
//ff:what TestParseSitemap_FileAndDirExclusion — 파일 파싱 + ParseDir 의 sitemap.html 제외 검증

package stml

import (
	"os"
	"path/filepath"
	"testing"
)

func TestParseSitemap_FileAndDirExclusion(t *testing.T) {
	dir := t.TempDir()
	sitemap := `<nav data-sitemap><ul><li data-page="login">로그인</li></ul></nav>`
	if err := os.WriteFile(filepath.Join(dir, "sitemap.html"), []byte(sitemap), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, "login.html"), []byte("<main></main>"), 0o644); err != nil {
		t.Fatal(err)
	}

	spec, diags := ParseSitemap(filepath.Join(dir, "sitemap.html"))
	if len(diags) != 0 {
		t.Fatalf("ParseSitemap diags = %+v", diags)
	}
	if len(spec.Navs) != 1 || spec.Navs[0].Items[0].Page != "login" {
		t.Errorf("spec = %+v", spec)
	}

	pages, diags := ParseDir(dir)
	if len(diags) != 0 {
		t.Fatalf("ParseDir diags = %+v", diags)
	}
	if len(pages) != 1 || pages[0].Name != "login" {
		t.Errorf("ParseDir must exclude sitemap.html, got pages %+v", pages)
	}
}
