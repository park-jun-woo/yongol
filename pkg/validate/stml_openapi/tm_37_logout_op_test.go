//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what TM-37 — 없는 op/GET op 발화, POST op·값 없는 logout·logout 부재 침묵 검증

package stml_openapi

import (
	"strings"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestTM37LogoutOp(t *testing.T) {
	opMap := buildOperationMethodMap(makeDoc(map[string]*openapi3.PathItem{
		"/auth/logout": postOp("Logout", nil),
		"/auth/me":     getOp("GetMe", nil, nil),
	}))
	layoutOf := func(logout *stml.LogoutSpec) []stml.LayoutSpec {
		return []stml.LayoutSpec{{Name: "app", File: "layouts/app.html", Logout: logout}}
	}

	t.Run("declared op is silent", func(t *testing.T) {
		diags := tm37LogoutOp(layoutOf(&stml.LogoutSpec{OperationID: "Logout"}), opMap)
		if len(diags) != 0 {
			t.Errorf("expected silence, got %+v", diags)
		}
	})

	t.Run("valueless logout is silent", func(t *testing.T) {
		diags := tm37LogoutOp(layoutOf(&stml.LogoutSpec{}), opMap)
		if len(diags) != 0 {
			t.Errorf("expected silence for valueless data-logout, got %+v", diags)
		}
	})

	t.Run("no logout is silent", func(t *testing.T) {
		diags := tm37LogoutOp(layoutOf(nil), opMap)
		if len(diags) != 0 {
			t.Errorf("expected silence without data-logout, got %+v", diags)
		}
	})

	t.Run("unknown operationId is an error", func(t *testing.T) {
		diags := tm37LogoutOp(layoutOf(&stml.LogoutSpec{OperationID: "SignOut"}), opMap)
		if got := countDiag(diags, "[TM-37]"); got != 1 {
			t.Fatalf("expected 1 TM-37, got %d: %+v", got, diags)
		}
		if diags[0].Level != diagnostic.LevelError {
			t.Errorf("Level = %v, want LevelError", diags[0].Level)
		}
		if diags[0].OperationID != "SignOut" {
			t.Errorf("OperationID = %q, want %q", diags[0].OperationID, "SignOut")
		}
		if !strings.Contains(diags[0].Message, "is not defined in OpenAPI") {
			t.Errorf("Message = %q, want not-defined branch", diags[0].Message)
		}
	})

	t.Run("GET op is an error", func(t *testing.T) {
		diags := tm37LogoutOp(layoutOf(&stml.LogoutSpec{OperationID: "GetMe"}), opMap)
		if got := countDiag(diags, "[TM-37]"); got != 1 {
			t.Fatalf("expected 1 TM-37, got %d: %+v", got, diags)
		}
		if !strings.Contains(diags[0].Message, "references a GET endpoint") {
			t.Errorf("Message = %q, want GET branch", diags[0].Message)
		}
	})
}
