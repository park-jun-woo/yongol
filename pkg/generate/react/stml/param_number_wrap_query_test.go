//ff:func feature=stml-gen type=test control=sequence
//ff:what integer path param이 useQuery에서 Number()로 래핑되는지 검증
package stml

import (
	"strings"
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestParamNumberWrapQuery(t *testing.T) {
	page, _ := stmlparser.ParseReader("building-detail.html", strings.NewReader(`<main>
  <article data-fetch="GetBuilding" data-param-id="route.id">
    <h1 data-bind="name"></h1>
  </article>
</main>`))
	code := GeneratePage(page, "", GenerateOptions{
		PathParamTypes: map[string]map[string]string{
			"GetBuilding": {"id": "integer"},
		},
	})
	assertContains(t, code, `id: Number(id)`)
	assertNotContains(t, code, `id: id }`)
}
