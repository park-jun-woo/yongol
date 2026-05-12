//ff:func feature=stml-gen type=test control=sequence
//ff:what data-each에서 스키마에 id 없으면 key={index} fallback 검증
package stml

import (
	"strings"
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

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
	assertContains(t, code, "key={index}")
	assertNotContains(t, code, "key={item.id}")
}
