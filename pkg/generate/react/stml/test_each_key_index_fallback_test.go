//ff:func feature=stml-gen type=test control=sequence
//ff:what data-each에서 스키마 없으면 key={index} fallback 검증
package stml

import (
	"strings"
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestEachKeyIndexFallback(t *testing.T) {
	page, _ := stmlparser.ParseReader("list-page.html", strings.NewReader(`<main>
  <section data-fetch="ListWorkflows">
    <ul data-each="items">
      <li>
        <span data-bind="title"></span>
      </li>
    </ul>
  </section>
</main>`))
	code := GeneratePage(page, "")
	assertContains(t, code, "key={index}")
	assertContains(t, code, ".map((item, index) =>")
}
