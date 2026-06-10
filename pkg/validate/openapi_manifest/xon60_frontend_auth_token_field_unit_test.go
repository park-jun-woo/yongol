//ff:func feature=validate type=test control=sequence topic=config-check
//ff:what xon60FrontendAuthTokenField — nil 가드 + token_field/refresh_field 실재/부재 검증
package openapi_manifest

import (
	"strings"
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXon60FrontendAuthTokenField_Unit(t *testing.T) {
	loginFields := map[string][]string{"Login": {"access_token", "refresh_token"}}

	t.Run("nil manifest returns nil", func(t *testing.T) {
		if diags := xon60FrontendAuthTokenField(&yongol.Fullstack{}); len(diags) != 0 {
			t.Fatalf("expected nil, got %+v", diags)
		}
	})

	t.Run("no frontend.auth returns nil", func(t *testing.T) {
		fs := xon60Fixture(nil, loginFields)
		if diags := xon60FrontendAuthTokenField(fs); len(diags) != 0 {
			t.Fatalf("expected nil, got %+v", diags)
		}
	})

	t.Run("nil OpenAPI doc returns nil", func(t *testing.T) {
		fs := &yongol.Fullstack{
			Manifest: &pmanifest.ProjectConfig{
				Frontend: pmanifest.Frontend{Auth: &pmanifest.FrontendAuth{TokenField: "access_token"}},
			},
		}
		if diags := xon60FrontendAuthTokenField(fs); len(diags) != 0 {
			t.Fatalf("expected nil, got %+v", diags)
		}
	})

	t.Run("token_field present passes", func(t *testing.T) {
		fs := xon60Fixture(&pmanifest.FrontendAuth{TokenField: "access_token"}, loginFields)
		if diags := xon60FrontendAuthTokenField(fs); len(diags) != 0 {
			t.Fatalf("expected 0, got %d: %+v", len(diags), diags)
		}
	})

	t.Run("token_field absent raises ERROR with candidates", func(t *testing.T) {
		fs := xon60Fixture(&pmanifest.FrontendAuth{TokenField: "acces_token"}, loginFields)
		diags := xon60FrontendAuthTokenField(fs)
		if len(diags) != 1 {
			t.Fatalf("expected 1, got %d: %+v", len(diags), diags)
		}
		if !strings.Contains(diags[0].Message, "[XON-60]") {
			t.Errorf("Message missing XON-60: %s", diags[0].Message)
		}
		if !strings.Contains(diags[0].Advice, "access_token") {
			t.Errorf("Advice missing candidate access_token: %s", diags[0].Advice)
		}
	})

	t.Run("refresh_field absent raises ERROR", func(t *testing.T) {
		fs := xon60Fixture(
			&pmanifest.FrontendAuth{TokenField: "access_token", RefreshField: "refressh_token"},
			loginFields,
		)
		diags := xon60FrontendAuthTokenField(fs)
		if len(diags) != 1 {
			t.Fatalf("expected 1, got %d: %+v", len(diags), diags)
		}
		if !strings.Contains(diags[0].Message, "refresh_field") {
			t.Errorf("Message should mention refresh_field: %s", diags[0].Message)
		}
	})

	t.Run("refresh_field present passes", func(t *testing.T) {
		fs := xon60Fixture(
			&pmanifest.FrontendAuth{TokenField: "access_token", RefreshField: "refresh_token"},
			loginFields,
		)
		if diags := xon60FrontendAuthTokenField(fs); len(diags) != 0 {
			t.Fatalf("expected 0, got %d: %+v", len(diags), diags)
		}
	})
}
