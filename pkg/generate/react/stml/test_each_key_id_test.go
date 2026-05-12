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
	// map callback should not include index parameter when id key is used
	assertContains(t, code, ".map((item) =>")
}

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
	// No ResponseArrayItemFields → fallback to index
	code := GeneratePage(page, "")
	assertContains(t, code, "key={index}")
	assertContains(t, code, ".map((item, index) =>")
}

func TestEachKeyIndexWhenNoIDInSchema(t *testing.T) {
	page, _ := stmlparser.ParseReader("list-page.html", strings.NewReader(`<main>
  <section data-fetch="ListLogs">
    <ul data-each="logs">
      <li>
        <span data-bind="message"></span>
      </li>
    </ul>
  </section>
</main>`))
	opt := GenerateOptions{
		ResponseArrayItemFields: map[string]map[string]map[string]bool{
			"ListLogs": {
				"logs": {"message": true, "timestamp": true},
			},
		},
	}
	code := GeneratePage(page, "", opt)
	// No "id" field in schema → fallback to index
	assertContains(t, code, "key={index}")
	assertNotContains(t, code, "key={item.id}")
}
