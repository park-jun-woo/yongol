//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what firstProtectedOpPage — 보호 op 소비 첫 페이지 반환 / 미존재 시 false 검증

package stml_openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestFirstProtectedOpPage(t *testing.T) {
	sec := openapi3.SecurityRequirements{openapi3.SecurityRequirement{"bearerAuth": {}}}
	doc := makeDoc(map[string]*openapi3.PathItem{
		"/public": {Get: &openapi3.Operation{OperationID: "GetPublic"}},
		"/secret": {Post: &openapi3.Operation{OperationID: "DoSecret", Security: &sec}},
	})
	opMap := buildOperationMethodMap(doc)

	t.Run("returns first page consuming a protected op", func(t *testing.T) {
		pages := []stml.PageSpec{
			{
				FileName: "public.html",
				Fetches: []stml.FetchBlock{
					{OperationID: "Unknown"},   // not in opMap → skipped
					{OperationID: "GetPublic"}, // unprotected → skipped
				},
			},
			{
				FileName: "secure.html",
				Actions:  []stml.ActionBlock{{OperationID: "DoSecret"}},
			},
		}
		fs := makeFS(pages, doc)
		got, ok := firstProtectedOpPage(fs, opMap)
		if !ok || got != "secure.html" {
			t.Errorf("expected (secure.html, true), got (%q, %v)", got, ok)
		}
	})

	t.Run("no protected op returns false", func(t *testing.T) {
		pages := []stml.PageSpec{
			{
				FileName: "public.html",
				Fetches:  []stml.FetchBlock{{OperationID: "GetPublic"}},
			},
		}
		fs := makeFS(pages, doc)
		got, ok := firstProtectedOpPage(fs, opMap)
		if ok || got != "" {
			t.Errorf("expected (\"\", false), got (%q, %v)", got, ok)
		}
	})
}
