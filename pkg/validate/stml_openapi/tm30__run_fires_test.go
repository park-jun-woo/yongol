//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what Run 경유 TM-30 — 응답 item 스키마에 없는 item.* 필드가 ERROR 로 발화함을 검증

package stml_openapi

import (
	"strings"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestTM30RunFires(t *testing.T) {
	photoItem := &openapi3.SchemaRef{Value: &openapi3.Schema{
		Type: &openapi3.Types{"object"},
		Properties: openapi3.Schemas{
			"id":      &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{"integer"}}},
			"caption": &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{"string"}}},
		},
	}}
	photos := &openapi3.SchemaRef{Value: &openapi3.Schema{
		Type:  &openapi3.Types{"array"},
		Items: photoItem,
	}}
	doc := makeDoc(map[string]*openapi3.PathItem{
		"/unit":   getOp("GetUnit", nil, map[string]*openapi3.SchemaRef{"photos": photos}),
		"/photos": {Delete: &openapi3.Operation{OperationID: "DeletePhoto"}},
	})

	page, parseDiags := stml.ParseReader("unit-info.html", strings.NewReader(`<main>
  <section data-fetch="GetUnit">
    <ul data-each="photos">
      <li>
        <span data-bind="caption"></span>
        <button data-action="DeletePhoto" data-param-photo-id="item.nope">x</button>
      </li>
    </ul>
  </section>
</main>`))
	if len(parseDiags) > 0 {
		t.Fatalf("unexpected parse diags: %v", parseDiags)
	}

	diags := Run(makeFS([]stml.PageSpec{page}, doc))
	if got := countDiag(diags, "[TM-30]"); got != 1 {
		t.Errorf("expected 1 TM-30 via Run, got %d: %+v", got, diags)
	}
}
