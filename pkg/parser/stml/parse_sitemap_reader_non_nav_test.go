//ff:func feature=stml-parse type=test control=sequence
//ff:what TestParseSitemapReader_TopLevelNonNavIsError — nav 가 아닌 최상위 요소는 파싱 에러 검증

package stml

import (
	"strings"
	"testing"
)

func TestParseSitemapReader_TopLevelNonNavIsError(t *testing.T) {
	src := `<div><ul><li data-page="login">로그인</li></ul></div>`
	_, diags := ParseSitemapReader("sitemap.html", strings.NewReader(src))
	if len(diags) != 1 || !strings.Contains(diags[0].Message, "<div>") {
		t.Fatalf("expected one top-level <div> error, got %+v", diags)
	}
}
