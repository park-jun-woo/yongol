//ff:func feature=stml-gen type=test control=sequence
//ff:what TestEachTypeAwareCell — data-each 셀의 img/boolean 타입 인지 방출과 RowLink 래핑 공존(BUG-126) 검증

package stml

import (
	"strings"
	"testing"

	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestEachTypeAwareCell(t *testing.T) {
	page, _ := stmlparser.ParseReader("building-detail.html", strings.NewReader(`<main>
  <section data-fetch="GetBuilding">
    <ul data-each="photos">
      <li>
        <img data-bind="url" />
        <span data-bind="featured"></span>
        <span data-bind="category"></span>
      </li>
    </ul>
  </section>
</main>`))
	opt := GenerateOptions{
		ResponseBindTypes: map[string]map[string]oapiparser.FieldTypeInfo{
			"GetBuilding": {
				"photos.url":      {Type: "string", Format: "uri"},
				"photos.featured": {Type: "boolean"},
				"photos.category": {Type: "string"},
			},
		},
	}
	code := GeneratePage(page, "", opt)

	// img column → <img src> cell, not a text cell
	assertContains(t, code, `<TD><img src={item.url} alt="Url" /></TD>`)
	assertNotContains(t, code, "<TD>{item.url}</TD>")
	// boolean column → Yes/No
	assertContains(t, code, "<TD>{item.featured ? 'Yes' : 'No'}</TD>")
	// plain string column unchanged
	assertContains(t, code, "<TD>{item.category}</TD>")
	// header still lists the field labels
	assertContains(t, code, "<TH>Url</TH>")

	// Whole-row <Link> still wraps the type-aware img cell content.
	linkPage, _ := stmlparser.ParseReader("gallery.html", strings.NewReader(`<main>
  <section data-fetch="GetBuilding">
    <ul data-each="photos">
      <li data-link="photo-detail" data-link-params="id:item.id">
        <img data-bind="url" />
      </li>
    </ul>
  </section>
</main>`))
	linkOpt := GenerateOptions{
		ResponseBindTypes: map[string]map[string]oapiparser.FieldTypeInfo{
			"GetBuilding": {"photos.url": {Type: "string"}},
		},
		RoutePatterns: map[string]string{"photo-detail": "/photos/:id"},
	}
	linkCode := GeneratePage(linkPage, "", linkOpt)
	assertContains(t, linkCode, "<TD><Link ")
	assertContains(t, linkCode, "<img src={item.url} alt=\"Url\" /></Link></TD>")
}
