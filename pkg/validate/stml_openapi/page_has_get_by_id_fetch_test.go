//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what pageHasGetByIdFetch — route param 소비 GET fetch 검출, non-GET·param 없음·미지 op 침묵 검증

package stml_openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestPageHasGetByIdFetch(t *testing.T) {
	opMap := buildOperationMethodMap(tm54Doc())

	withRouteParam := stml.PageSpec{Fetches: []stml.FetchBlock{
		{OperationID: "GetRule", Params: []stml.ParamBind{{Name: "RuleID", Source: "route.RuleID"}}},
	}}
	if !pageHasGetByIdFetch(withRouteParam, opMap) {
		t.Errorf("GET with a route param should be detected")
	}

	noParam := stml.PageSpec{Fetches: []stml.FetchBlock{{OperationID: "GetRule"}}}
	if pageHasGetByIdFetch(noParam, opMap) {
		t.Errorf("GET without a route param is not GET-by-id")
	}

	itemParam := stml.PageSpec{Fetches: []stml.FetchBlock{
		{OperationID: "GetRule", Params: []stml.ParamBind{{Name: "id", Source: "item.id"}}},
	}}
	if pageHasGetByIdFetch(itemParam, opMap) {
		t.Errorf("item.* source is not a route path param")
	}

	// PUT op consumed as a fetch (unusual) — not GET, ignored.
	putFetch := stml.PageSpec{Fetches: []stml.FetchBlock{
		{OperationID: "UpdateRule", Params: []stml.ParamBind{{Name: "RuleID", Source: "route.RuleID"}}},
	}}
	if pageHasGetByIdFetch(putFetch, opMap) {
		t.Errorf("non-GET method should not count")
	}

	// Unknown op.
	unknown := stml.PageSpec{Fetches: []stml.FetchBlock{
		{OperationID: "NoSuch", Params: []stml.ParamBind{{Name: "RuleID", Source: "route.RuleID"}}},
	}}
	if pageHasGetByIdFetch(unknown, opMap) {
		t.Errorf("unknown op should be ignored")
	}

	_ = openapi3.T{}
}
