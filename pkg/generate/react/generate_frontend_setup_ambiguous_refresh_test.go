//ff:func feature=gen-react type=test control=sequence
//ff:what generateFrontendSetup — refresh op 추론 모호(후보 2+) 시 generate ERROR 로 조기 중단 검증

package react

import (
	"strings"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestGenerateFrontendSetup_AmbiguousRefreshOp_Error(t *testing.T) {
	tokenResp := []string{"access_token", "refresh_token"}
	fs := &yongol.Fullstack{
		Manifest: &manifest.ProjectConfig{
			Backend: manifest.Backend{Auth: &manifest.Auth{Mode: "bearer"}},
			Frontend: manifest.Frontend{Auth: &manifest.FrontendAuth{
				TokenField: "access_token", RefreshField: "refresh_token",
			}},
		},
		// Two token-yielding ops, neither bound by a data-capture
		// declaration -> structural inference cannot pick one.
		OpenAPIDoc: &openapi3.T{Paths: openapi3.NewPaths(
			openapi3.WithPath("/auth/login", &openapi3.PathItem{Post: buildTokenOp("Login", tokenResp, []string{"email"})}),
			openapi3.WithPath("/auth/refresh", &openapi3.PathItem{Post: buildTokenOp("Refresh", tokenResp, []string{"refresh_token"})}),
		)},
	}

	err := generateFrontendSetup(fs, t.TempDir())
	if err == nil {
		t.Fatal("ambiguous refresh-op inference: want generate error, got nil")
	}
	if !strings.Contains(err.Error(), "refresh_op") {
		t.Errorf("error should advise declaring frontend.auth.refresh_op, got: %v", err)
	}
}
