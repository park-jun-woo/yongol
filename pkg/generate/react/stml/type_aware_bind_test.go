//ff:func feature=stml-gen type=test control=sequence
//ff:what TestTypeAwareBind — boolean/img/number/date-time data-bind이 타입별 JSX로 방출되는지(BUG-126) 검증

package stml

import (
	"strings"
	"testing"

	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestTypeAwareBind(t *testing.T) {
	page, _ := stmlparser.ParseReader("building-detail.html", strings.NewReader(`<main>
  <section data-fetch="GetBuilding">
    <span data-bind="can_delete"></span>
    <span data-bind="created_at"></span>
    <span data-bind="credits"></span>
    <img data-bind="thumbnail" />
    <span data-bind="name"></span>
  </section>
</main>`))
	opt := GenerateOptions{
		ResponseBindTypes: map[string]map[string]oapiparser.FieldTypeInfo{
			"GetBuilding": {
				"can_delete": {Type: "boolean"},
				"created_at": {Type: "string", Format: "date-time"},
				"credits":    {Type: "integer"},
				"thumbnail":  {Type: "string"},
				"name":       {Type: "string"},
			},
		},
	}
	code := GeneratePage(page, "", opt)

	// boolean → Yes/No (no longer blank)
	assertContains(t, code, "{getBuildingData.can_delete ? 'Yes' : 'No'}")
	// date-time → locale format
	assertContains(t, code, "{getBuildingData.created_at ? new Date(getBuildingData.created_at).toLocaleString() : ''}")
	// integer → toLocaleString
	assertContains(t, code, "{typeof getBuildingData.credits === 'number' ? getBuildingData.credits.toLocaleString() : getBuildingData.credits}")
	// <img> → src binding, self-closing, no text children
	assertContains(t, code, `<img src={getBuildingData.thumbnail} alt="Thumbnail" />`)
	assertNotContains(t, code, "<img>{getBuildingData.thumbnail}</img>")
	// plain string field stays {value}
	assertContains(t, code, "<span>{getBuildingData.name}</span>")

	// Without ResponseBindTypes the emission is byte-identical to the legacy
	// plain-children form (regression: boolean is NOT rewritten when unwired).
	plain := GeneratePage(page, "", GenerateOptions{})
	assertContains(t, plain, "<span>{getBuildingData.can_delete}</span>")
	assertContains(t, plain, "<span>{getBuildingData.name}</span>")
}
