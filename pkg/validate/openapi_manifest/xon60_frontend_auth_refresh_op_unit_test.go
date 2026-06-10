//ff:func feature=validate type=test control=sequence topic=config-check
//ff:what xon60FrontendAuthTokenField — refresh_op 미존재/token_field 부재/정상 변형 검증
package openapi_manifest

import (
	"strings"
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

func TestXon60FrontendAuthRefreshOp_Unit(t *testing.T) {
	loginFields := map[string][]string{"Login": {"access_token", "refresh_token"}}

	t.Run("refresh_op unknown operationId raises ERROR", func(t *testing.T) {
		fs := xon60Fixture(
			&pmanifest.FrontendAuth{TokenField: "access_token", RefreshOp: "RefreshToken"},
			loginFields,
		)
		diags := xon60FrontendAuthTokenField(fs)
		if len(diags) != 1 {
			t.Fatalf("expected 1, got %d: %+v", len(diags), diags)
		}
		if !strings.Contains(diags[0].Message, "RefreshToken") {
			t.Errorf("Message should mention RefreshToken: %s", diags[0].Message)
		}
		if !strings.Contains(diags[0].Advice, "Login") {
			t.Errorf("Advice should list candidate operationIds: %s", diags[0].Advice)
		}
	})

	t.Run("refresh_op without token_field in response raises ERROR", func(t *testing.T) {
		fs := xon60Fixture(
			&pmanifest.FrontendAuth{TokenField: "access_token", RefreshOp: "RefreshToken"},
			map[string][]string{
				"Login":        {"access_token", "refresh_token"},
				"RefreshToken": {"expires_at"},
			},
		)
		diags := xon60FrontendAuthTokenField(fs)
		if len(diags) != 1 {
			t.Fatalf("expected 1, got %d: %+v", len(diags), diags)
		}
		if !strings.Contains(diags[0].Advice, "expires_at") {
			t.Errorf("Advice should list the refresh op's fields: %s", diags[0].Advice)
		}
	})

	t.Run("refresh_op with token_field passes", func(t *testing.T) {
		fs := xon60Fixture(
			&pmanifest.FrontendAuth{TokenField: "access_token", RefreshOp: "RefreshToken"},
			map[string][]string{
				"Login":        {"access_token", "refresh_token"},
				"RefreshToken": {"access_token", "refresh_token"},
			},
		)
		if diags := xon60FrontendAuthTokenField(fs); len(diags) != 0 {
			t.Fatalf("expected 0, got %d: %+v", len(diags), diags)
		}
	})
}
