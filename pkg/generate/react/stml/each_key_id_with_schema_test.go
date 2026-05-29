//ff:func feature=stml-gen type=test control=sequence
//ff:what data-each에서 KeyField 설정 시 key={item.id} 생성을 검증
package stml

import (
	"strings"
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestEachKeyID(t *testing.T) {
	page, _ := stmlparser.ParseReader("list-page.html", strings.NewReader(`<main>
  <section data-fetch="ListWorkflows">
    <ul data-each="items">
      <li>
        <span data-bind="title"></span>
        <span data-bind="status"></span>
      </li>
    </ul>
  </section>
</main>`))
	opt := GenerateOptions{
		ResponseArrayItemFields: map[string]map[string]map[string]bool{
			"ListWorkflows": {
				"items": {"id": true, "title": true, "status": true},
			},
		},
	}
	code := GeneratePage(page, "", opt)
	assertContains(t, code, "key={item.id}")
	assertNotContains(t, code, "key={index}")
	assertContains(t, code, ".map((item) =>")
}
