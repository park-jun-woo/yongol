//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what Run 경유 TM-26/TM-33 — 페이지명 redirect 의 없는 respField 가 ERROR 로 발화함을 검증

package stml_openapi

import (
	"strings"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestTM33RunFires(t *testing.T) {
	doc := makeDoc(map[string]*openapi3.PathItem{
		"/contracts": postOpWithResp("CreateContract", map[string]*openapi3.SchemaRef{
			"id": {Value: &openapi3.Schema{Type: &openapi3.Types{"integer"}}},
		}),
	})

	// The substituted respField is not in the 2xx response schema → TM-33;
	// the page-name target itself exists, so TM-26 stays silent.
	page, parseDiags := stml.ParseReader("contract-new.html", strings.NewReader(`<main>
  <div data-action="CreateContract" data-redirect="contract-edit" data-redirect-params="contract_id -> ContractID">
    <input data-field="Title" type="text" />
    <button type="submit">등록</button>
  </div>
</main>`))
	if len(parseDiags) > 0 {
		t.Fatalf("unexpected parse diags: %v", parseDiags)
	}
	edit := stml.PageSpec{
		Name:     "contract-edit",
		FileName: "contract-edit.html",
		Route:    "/contract-edit/:ContractID",
	}

	diags := Run(makeFS([]stml.PageSpec{page, edit}, doc))
	if got := countDiag(diags, "[TM-33]"); got != 1 {
		t.Errorf("expected 1 TM-33 via Run, got %d: %+v", got, diags)
	}
	if got := countDiag(diags, "[TM-26]"); got != 0 {
		t.Errorf("expected 0 TM-26 via Run, got %d: %+v", got, diags)
	}
}
