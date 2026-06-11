//ff:func feature=stml-parse type=test control=sequence
//ff:what TestParseSitemapReader_PageAndHrefBothKept — data-page 와 <a href> 동시 보유 시 둘 다 보존(TM-39 가 거부) 검증

package stml

import (
	"strings"
	"testing"
)

func TestParseSitemapReader_PageAndHrefBothKept(t *testing.T) {
	src := `<nav data-sitemap><ul><li data-page="docs"><a href="https://docs.example.com">문서</a></li></ul></nav>`
	spec, diags := ParseSitemapReader("sitemap.html", strings.NewReader(src))
	if len(diags) != 0 {
		t.Fatalf("expected no diags, got %+v", diags)
	}
	node := spec.Navs[0].Items[0]
	// The parser preserves the contradiction — TM-39 rejects it.
	if node.Page != "docs" || node.Href != "https://docs.example.com" {
		t.Errorf("node = %+v, want both page and href kept", node)
	}
}
