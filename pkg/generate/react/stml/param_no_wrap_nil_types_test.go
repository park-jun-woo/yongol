//ff:func feature=stml-gen type=test control=sequence
//ff:what PathParamTypes가 nil이면 Number() 래핑 없이 기존 동작을 유지하는지 검증
package stml

import (
	"strings"
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestParamNoWrapNilTypes(t *testing.T) {
	page, _ := stmlparser.ParseReader("building-detail.html", strings.NewReader(`<main>
  <article data-fetch="GetBuilding" data-param-id="route.id">
    <h1 data-bind="name"></h1>
  </article>
</main>`))
	code := GeneratePage(page, "", GenerateOptions{})
	assertContains(t, code, `id: id`)
	assertNotContains(t, code, `Number(id)`)
}
