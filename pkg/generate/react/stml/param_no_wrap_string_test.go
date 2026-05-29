//ff:func feature=stml-gen type=test control=sequence
//ff:what string path param에는 Number() 래핑이 적용되지 않는지 검증
package stml

import (
	"strings"
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestParamNoWrapString(t *testing.T) {
	page, _ := stmlparser.ParseReader("user-detail.html", strings.NewReader(`<main>
  <article data-fetch="GetUser" data-param-slug="route.slug">
    <h1 data-bind="name"></h1>
  </article>
</main>`))
	code := GeneratePage(page, "", GenerateOptions{
		PathParamTypes: map[string]map[string]string{
			"GetUser": {"slug": "string"},
		},
	})
	assertContains(t, code, `slug: slug`)
	assertNotContains(t, code, `Number(slug)`)
}
