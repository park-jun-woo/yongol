//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what Run 경유 TM-31/TM-32 — 오타 페이지명·필수 세그먼트 누락이 ERROR 로 발화함을 검증

package stml_openapi

import (
	"strings"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestTM31TM32RunFires(t *testing.T) {
	buildingItem := &openapi3.SchemaRef{Value: &openapi3.Schema{
		Type: &openapi3.Types{"object"},
		Properties: openapi3.Schemas{
			"id":   &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{"integer"}}},
			"name": &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{"string"}}},
		},
	}}
	buildings := &openapi3.SchemaRef{Value: &openapi3.Schema{
		Type:  &openapi3.Types{"array"},
		Items: buildingItem,
	}}
	doc := makeDoc(map[string]*openapi3.PathItem{
		"/buildings": getOp("ListBuildings", nil, map[string]*openapi3.SchemaRef{"buildings": buildings}),
	})

	// Typo'd target → TM-31; row link to the real detail page without the
	// required segment mapping → TM-32.
	page, parseDiags := stml.ParseReader("building-list.html", strings.NewReader(`<main>
  <section data-fetch="ListBuildings">
    <ul data-each="buildings">
      <li data-link="building-detail"><span data-bind="name"></span></li>
    </ul>
  </section>
  <a data-link="bulding-detail">상세 (오타)</a>
</main>`))
	if len(parseDiags) > 0 {
		t.Fatalf("unexpected parse diags: %v", parseDiags)
	}
	detail := stml.PageSpec{
		Name:     "building-detail",
		FileName: "building-detail.html",
		Route:    "/buildings/:BuildingID",
	}

	diags := Run(makeFS([]stml.PageSpec{page, detail}, doc))
	if got := countDiag(diags, "[TM-31]"); got != 1 {
		t.Errorf("expected 1 TM-31 via Run, got %d: %+v", got, diags)
	}
	if got := countDiag(diags, "[TM-32]"); got != 1 {
		t.Errorf("expected 1 TM-32 via Run, got %d: %+v", got, diags)
	}
}
