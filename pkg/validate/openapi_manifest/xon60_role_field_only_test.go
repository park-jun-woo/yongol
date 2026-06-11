//ff:func feature=validate type=test control=sequence topic=config-check
//ff:what xon60 role_field 전용 블록 면제 + 혼재 블록·bearer token_field 누락 ERROR 불변 회귀 검증

package openapi_manifest

import (
	"strings"
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

func TestXon60RoleFieldOnlyExemption(t *testing.T) {
	loginFields := map[string][]string{"Login": {"access_token", "refresh_token"}}

	t.Run("role_field-only block is exempt from token_field", func(t *testing.T) {
		// Cookie-mode data-roles wiring (plans/stml/sitemap Phase005): the
		// block declares no token contract, so nothing is verified here.
		fs := xon60Fixture(&pmanifest.FrontendAuth{RoleField: "role"}, loginFields)
		if diags := xon60FrontendAuthTokenField(fs); len(diags) != 0 {
			t.Fatalf("expected 0, got %d: %+v", len(diags), diags)
		}
	})

	t.Run("role_field beside token keys keeps the full check", func(t *testing.T) {
		// Mixed block: any token-related key restores the token_field
		// requirement — XON-60 is its single enforcer.
		fs := xon60Fixture(&pmanifest.FrontendAuth{RoleField: "role", RefreshField: "refresh_token"}, loginFields)
		diags := xon60FrontendAuthTokenField(fs)
		if len(diags) != 1 {
			t.Fatalf("expected 1, got %d: %+v", len(diags), diags)
		}
		if !strings.Contains(diags[0].Message, "token_field") {
			t.Errorf("Message should report the missing token_field: %s", diags[0].Message)
		}
	})

	t.Run("bearer regression: missing token_field stays ERROR", func(t *testing.T) {
		// The Phase005 exemption must not weaken the bearer protection: a
		// frontend.auth block without role_field and without token_field
		// still fails exactly as before.
		fs := xon60Fixture(&pmanifest.FrontendAuth{Store: "memory"}, loginFields)
		diags := xon60FrontendAuthTokenField(fs)
		if len(diags) != 1 {
			t.Fatalf("expected 1, got %d: %+v", len(diags), diags)
		}
		if !strings.Contains(diags[0].Message, "[XON-60]") || !strings.Contains(diags[0].Message, "token_field") {
			t.Errorf("Message should be the XON-60 token_field ERROR: %s", diags[0].Message)
		}
	})
}
