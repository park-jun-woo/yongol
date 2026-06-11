//ff:func feature=validate type=test-helper control=sequence topic=stml-openapi
//ff:what tm56Page — PatchRule 폼(필드 2개) STML 페이지 파싱

package stml_openapi

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// tm56Page parses a page with a single PATCH PatchRule form carrying two fields.
func tm56Page(t *testing.T) stml.PageSpec {
	t.Helper()
	src := `<main>
  <form data-action="PatchRule" data-param-rule-id="route.RuleID">
    <input data-field="sheet_name" />
    <input data-field="start_row" />
    <button type="submit">save</button>
  </form>
</main>`
	page, diags := stml.ParseReader("rule-edit.html", strings.NewReader(src))
	if len(diags) > 0 {
		t.Fatal(diags)
	}
	return page
}
