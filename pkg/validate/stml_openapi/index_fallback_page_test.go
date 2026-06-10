//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what indexFallbackPage — 파일명 정렬 첫 공개 페이지 / 보호·param 스킵 / 전부 보호 시 /login 검증

package stml_openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestIndexFallbackPage(t *testing.T) {
	t.Run("first page by file-name sort", func(t *testing.T) {
		fs := makeFS([]stml.PageSpec{
			{Name: "settings", FileName: "settings.html"},
			{Name: "dashboard", FileName: "dashboard.html"},
		}, nil)
		file, path := indexFallbackPage(fs, nil)
		if file != "dashboard.html" || path != "/dashboard" {
			t.Errorf("got (%q, %q), want (dashboard.html, /dashboard)", file, path)
		}
	})

	t.Run("protected and parameterized pages are skipped", func(t *testing.T) {
		doc := makeDoc(map[string]*openapi3.PathItem{"/me": securedGetOp("GetMe")})
		fs := makeFS([]stml.PageSpec{
			{Name: "dashboard", FileName: "dashboard.html",
				Fetches: []stml.FetchBlock{{OperationID: "GetMe"}}},
			{Name: "item-detail", FileName: "item-detail.html",
				Fetches: []stml.FetchBlock{{OperationID: "GetItem",
					Params: []stml.ParamBind{{Name: "id", Source: "route.ItemID"}}}}},
			{Name: "settings", FileName: "settings.html"},
		}, doc)
		fs.Manifest.Backend.Auth = &manifest.Auth{Type: "jwt", Mode: "bearer"}
		file, path := indexFallbackPage(fs, buildOperationMethodMap(doc))
		if file != "settings.html" || path != "/settings" {
			t.Errorf("got (%q, %q), want (settings.html, /settings)", file, path)
		}
	})

	t.Run("every candidate excluded falls back to /login", func(t *testing.T) {
		doc := makeDoc(map[string]*openapi3.PathItem{"/me": securedGetOp("GetMe")})
		fs := makeFS([]stml.PageSpec{
			{Name: "dashboard", FileName: "dashboard.html",
				Fetches: []stml.FetchBlock{{OperationID: "GetMe"}}},
		}, doc)
		fs.Manifest.Backend.Auth = &manifest.Auth{Type: "jwt", Mode: "bearer"}
		file, path := indexFallbackPage(fs, buildOperationMethodMap(doc))
		if file != "" || path != "/login" {
			t.Errorf("got (%q, %q), want (\"\", /login)", file, path)
		}
	})

	t.Run("no backend.auth keeps protected-looking pages public", func(t *testing.T) {
		// Mirrors react.resolveProtectedPages returning nil without auth —
		// security on ops cannot gate routes that have no ProtectedRoute.
		doc := makeDoc(map[string]*openapi3.PathItem{"/me": securedGetOp("GetMe")})
		fs := makeFS([]stml.PageSpec{
			{Name: "dashboard", FileName: "dashboard.html",
				Fetches: []stml.FetchBlock{{OperationID: "GetMe"}}},
		}, doc)
		file, path := indexFallbackPage(fs, buildOperationMethodMap(doc))
		if file != "dashboard.html" || path != "/dashboard" {
			t.Errorf("got (%q, %q), want (dashboard.html, /dashboard)", file, path)
		}
	})
}
