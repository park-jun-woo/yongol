//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what TM-35 — 인덱스 미선언 WARNING 발화(폴백 페이지명 명시)와 선언/"/" 점유/프론트 OFF 침묵 검증

package stml_openapi

import (
	"strings"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestTM35IndexFallback(t *testing.T) {
	pages := []stml.PageSpec{
		{Name: "forgot-password", FileName: "forgot-password.html"},
		{Name: "login", FileName: "login.html"},
	}

	t.Run("undeclared index fires and names the fallback page", func(t *testing.T) {
		fs := makeFS(pages, nil)
		diags := tm35IndexFallback(fs, nil)
		if got := countDiag(diags, "[TM-35]"); got != 1 {
			t.Fatalf("expected 1 TM-35, got %d: %+v", got, diags)
		}
		if diags[0].Level != diagnostic.LevelWarning {
			t.Errorf("Level = %v, want LevelWarning", diags[0].Level)
		}
		// File-name sort picks forgot-password — the BUG-114 (3) accident
		// the warning must make visible.
		if !strings.Contains(diags[0].Message, "forgot-password.html") {
			t.Errorf("Message should name the fallback page, got %q", diags[0].Message)
		}
		if !strings.Contains(diags[0].Advice, "frontend.index") || !strings.Contains(diags[0].Advice, `data-route="/"`) {
			t.Errorf("Advice should name both declaration vehicles, got %q", diags[0].Advice)
		}
	})

	t.Run("all candidates protected falls back to /login", func(t *testing.T) {
		doc := makeDoc(map[string]*openapi3.PathItem{
			"/me": securedGetOp("GetMe"),
		})
		fs := makeFS([]stml.PageSpec{
			{
				Name:     "dashboard",
				FileName: "dashboard.html",
				Fetches:  []stml.FetchBlock{{OperationID: "GetMe"}},
			},
		}, doc)
		fs.Manifest.Backend.Auth = &manifest.Auth{Type: "jwt", Mode: "bearer"}
		diags := tm35IndexFallback(fs, buildOperationMethodMap(doc))
		if got := countDiag(diags, "[TM-35]"); got != 1 {
			t.Fatalf("expected 1 TM-35, got %d: %+v", got, diags)
		}
		if !strings.Contains(diags[0].Message, "/login") {
			t.Errorf("Message should name the /login fallback, got %q", diags[0].Message)
		}
	})

	t.Run("declared frontend.index is silent", func(t *testing.T) {
		fs := makeFS(pages, nil)
		fs.Manifest.Frontend.Index = "login"
		if diags := tm35IndexFallback(fs, nil); len(diags) != 0 {
			t.Errorf("expected silence with frontend.index declared, got %+v", diags)
		}
	})

	t.Run("slash mount is silent", func(t *testing.T) {
		fs := makeFS([]stml.PageSpec{
			{Name: "home", FileName: "home.html", Route: "/"},
			{Name: "login", FileName: "login.html"},
		}, nil)
		if diags := tm35IndexFallback(fs, nil); len(diags) != 0 {
			t.Errorf("expected silence when a page mounts /, got %+v", diags)
		}
	})

	t.Run("frontend off is silent", func(t *testing.T) {
		fs := makeFS(pages, nil)
		off := false
		fs.Manifest.Frontend.Enabled = &off
		if diags := tm35IndexFallback(fs, nil); len(diags) != 0 {
			t.Errorf("expected silence with frontend OFF, got %+v", diags)
		}
	})

	t.Run("zero pages is silent", func(t *testing.T) {
		fs := makeFS(nil, nil)
		if diags := tm35IndexFallback(fs, nil); len(diags) != 0 {
			t.Errorf("expected silence with zero pages (XMO-11 owns it), got %+v", diags)
		}
	})
}
