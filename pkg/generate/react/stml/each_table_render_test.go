//ff:func feature=stml-gen type=test control=sequence
//ff:what data-each가 Table 구조(THead/TBody/TR/TH/TD)로 렌더되는지 검증
package stml

import (
	"strings"
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestEachTableRender(t *testing.T) {
	page, _ := stmlparser.ParseReader("list-page.html", strings.NewReader(`<main>
  <section data-fetch="ListWorkflows">
    <ul data-each="workflows">
      <li>
        <span data-bind="title"></span>
        <span data-bind="action_type"></span>
        <span data-bind="created_at"></span>
      </li>
    </ul>
  </section>
</main>`))
	opt := GenerateOptions{
		ResponseArrayItemFields: map[string]map[string]map[string]bool{
			"ListWorkflows": {
				"workflows": {"id": true, "title": true, "action_type": true, "created_at": true},
			},
		},
	}
	code := GeneratePage(page, "", opt)

	// Table structure
	assertContains(t, code, "<Table>")
	assertContains(t, code, "</Table>")
	assertContains(t, code, "<THead>")
	assertContains(t, code, "</THead>")
	assertContains(t, code, "<TBody>")
	assertContains(t, code, "</TBody>")

	// Header labels (snake_case → Title Case)
	assertContains(t, code, "<TH>Title</TH>")
	assertContains(t, code, "<TH>Action Type</TH>")
	assertContains(t, code, "<TH>Created At</TH>")

	// Data cells
	assertContains(t, code, "<TD>{item.title}</TD>")
	assertContains(t, code, "<TD>{item.action_type}</TD>")
	assertContains(t, code, "<TD>{item.created_at}</TD>")

	// Key from schema (has id)
	assertContains(t, code, "key={item.id}")

	// Table import
	assertContains(t, code, "import { Table, THead, TBody, TR, TH, TD } from '@/components/ui/Table'")

	// No ul/li tags
	assertNotContains(t, code, "<ul")
	assertNotContains(t, code, "<li")
}
