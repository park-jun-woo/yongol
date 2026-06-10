//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what pageConsumesProtectedOp — 보호 op 소비 페이지 true / 공개 op·미등록 op 페이지 false 검증

package stml_openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestPageConsumesProtectedOp(t *testing.T) {
	doc := makeDoc(map[string]*openapi3.PathItem{
		"/me":    securedGetOp("GetMe"),
		"/items": getOp("ListItems", nil, nil),
	})
	fs := makeFS(nil, doc)
	opMap := buildOperationMethodMap(doc)
	ops := map[string]struct{}{"GetMe": {}, "ListItems": {}}

	protected := stml.PageSpec{Name: "dashboard", FileName: "dashboard.html",
		Fetches: []stml.FetchBlock{{OperationID: "GetMe"}}}
	if !pageConsumesProtectedOp(protected, fs, opMap, ops) {
		t.Error("expected true for a page consuming a secured op")
	}

	public := stml.PageSpec{Name: "items", FileName: "items.html",
		Fetches: []stml.FetchBlock{{OperationID: "ListItems"}}}
	if pageConsumesProtectedOp(public, fs, opMap, ops) {
		t.Error("expected false for a page consuming only public ops")
	}

	unknown := stml.PageSpec{Name: "ghost", FileName: "ghost.html",
		Fetches: []stml.FetchBlock{{OperationID: "NoSuchOp"}}}
	if pageConsumesProtectedOp(unknown, fs, opMap, ops) {
		t.Error("expected false for an op absent from the opMap")
	}
}
