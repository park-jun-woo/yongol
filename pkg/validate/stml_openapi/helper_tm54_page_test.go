//ff:func feature=validate type=test-helper control=sequence topic=stml-openapi
//ff:what tm54Page — GetRule fetch에 중첩된 UpdateRule edit 폼 STML 페이지 파싱 (prefill/fields 가변)

package stml_openapi

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// tm54Page parses a canonical edit page: a GET-by-id GetRule fetch wrapping an
// UpdateRule form. prefill is the data-prefill attribute snippet (e.g.
// ` data-prefill="GetRule"` or ""), fields the data-field inputs markup.
func tm54Page(t *testing.T, prefill string, fields string) stml.PageSpec {
	t.Helper()
	src := `<main>
  <article data-fetch="GetRule" data-param-rule-id="route.RuleID">
    <form data-action="UpdateRule"` + prefill + ` data-param-rule-id="route.RuleID">
` + fields + `
      <button type="submit">save</button>
    </form>
  </article>
</main>`
	page, diags := stml.ParseReader("rule-edit.html", strings.NewReader(src))
	if len(diags) > 0 {
		t.Fatal(diags)
	}
	return page
}
