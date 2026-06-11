//ff:func feature=stml-parse type=test control=sequence
//ff:what TestParseSitemapReader_EmptyFileIsError — nav 블록 없는 빈 sitemap 은 파싱 에러 검증

package stml

import (
	"strings"
	"testing"
)

func TestParseSitemapReader_EmptyFileIsError(t *testing.T) {
	_, diags := ParseSitemapReader("sitemap.html", strings.NewReader(""))
	if len(diags) != 1 || !strings.Contains(diags[0].Message, "no <nav data-sitemap>") {
		t.Fatalf("expected empty-sitemap error, got %+v", diags)
	}
}
