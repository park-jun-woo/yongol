//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what TM-58 — bearer+값 없는 data-logout+OpenAPI logout op 존재 시 WARNING, 나머지 침묵 검증

package stml_openapi

import (
	"strings"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestTM58BearerLogoutOpHint(t *testing.T) {
	fsWith := func(auth *manifest.Auth, logout *stml.LogoutSpec, doc *openapi3.T) *yongol.Fullstack {
		fs := makeFS(nil, doc)
		fs.Manifest.Backend.Auth = auth
		if logout != nil {
			fs.Layouts = []stml.LayoutSpec{{Name: "app", File: "layouts/app.html", Logout: logout}}
		}
		return fs
	}

	bearerAuth := &manifest.Auth{Type: "jwt", Mode: "bearer"}
	cookieAuth := &manifest.Auth{Type: "jwt", Mode: "cookie"}
	doc := logoutDoc("Logout")

	t.Run("bearer valueless data-logout with logout op fires warning", func(t *testing.T) {
		diags := tm58BearerLogoutOpHint(fsWith(bearerAuth, &stml.LogoutSpec{}, doc))
		if got := countDiag(diags, "[TM-58]"); got != 1 {
			t.Fatalf("expected 1 TM-58, got %d: %+v", got, diags)
		}
		if diags[0].Level != diagnostic.LevelWarning {
			t.Errorf("Level = %v, want LevelWarning", diags[0].Level)
		}
		if !strings.Contains(diags[0].Message, "operationId") {
			t.Errorf("Message = %q, want operationId mention", diags[0].Message)
		}
	})

	t.Run("bearer data-logout with operationId is silent", func(t *testing.T) {
		diags := tm58BearerLogoutOpHint(fsWith(bearerAuth, &stml.LogoutSpec{OperationID: "Logout"}, doc))
		if len(diags) != 0 {
			t.Errorf("expected silence, got %+v", diags)
		}
	})

	t.Run("cookie mode valueless data-logout is silent", func(t *testing.T) {
		diags := tm58BearerLogoutOpHint(fsWith(cookieAuth, &stml.LogoutSpec{}, doc))
		if len(diags) != 0 {
			t.Errorf("expected silence for cookie mode (TM-38 territory), got %+v", diags)
		}
	})

	t.Run("bearer valueless data-logout without logout op is silent", func(t *testing.T) {
		emptyDoc := &openapi3.T{Paths: openapi3.NewPaths()}
		diags := tm58BearerLogoutOpHint(fsWith(bearerAuth, &stml.LogoutSpec{}, emptyDoc))
		if len(diags) != 0 {
			t.Errorf("expected silence without logout op, got %+v", diags)
		}
	})

	t.Run("bearer no layout logout is silent", func(t *testing.T) {
		diags := tm58BearerLogoutOpHint(fsWith(bearerAuth, nil, doc))
		if len(diags) != 0 {
			t.Errorf("expected silence without data-logout, got %+v", diags)
		}
	})

	t.Run("no auth is silent", func(t *testing.T) {
		diags := tm58BearerLogoutOpHint(fsWith(nil, &stml.LogoutSpec{}, doc))
		if len(diags) != 0 {
			t.Errorf("expected silence without auth, got %+v", diags)
		}
	})
}
