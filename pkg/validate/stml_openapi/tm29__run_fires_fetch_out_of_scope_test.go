//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what Run 경유 TM-29 — 액션 블록에서 발화하고 data-fetch(GET)는 4xx 선언이어도 대상 외임을 검증

package stml_openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestTM29RunFiresAndFetchOutOfScope(t *testing.T) {
	desc := "error"
	errResp := &openapi3.ResponseRef{Value: &openapi3.Response{Description: &desc}}
	listItem := getOp("ListItems", nil, map[string]*openapi3.SchemaRef{"items": arrayProp("string")})
	listItem.Get.Responses.Set("500", errResp)
	createItem := postOp("CreateItem", nil)
	createItem.Post.Responses.Set("400", errResp)

	doc := makeDoc(map[string]*openapi3.PathItem{
		"/items":  listItem,
		"/create": createItem,
	})
	fs := makeFS([]stml.PageSpec{{
		Name:     "page",
		FileName: "page.html",
		Fetches:  []stml.FetchBlock{{OperationID: "ListItems"}},
		Actions:  []stml.ActionBlock{{OperationID: "CreateItem"}},
	}}, doc)

	diags := Run(fs)
	// The action block without data-on-error fires exactly once; the
	// data-fetch (GET) block is out of scope even though its operation
	// declares a 5xx response.
	if got := countDiag(diags, "[TM-29]"); got != 1 {
		t.Errorf("expected 1 TM-29 (action only, fetch out of scope), got %d: %+v", got, diags)
	}
}
