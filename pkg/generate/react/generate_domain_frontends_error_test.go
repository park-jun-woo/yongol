//ff:func feature=gen-react type=test control=sequence
//ff:what generateDomainFrontends — 한 도메인 view 가 generate ERROR(모호 refresh_op) 면 루프가 즉시 그 에러를 전파

package react

import (
	"strings"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestGenerateDomainFrontends_PropagatesDomainError(t *testing.T) {
	tokenResp := []string{"access_token", "refresh_token"}
	// A domain whose view yields two token-bearing ops with no capture
	// binding -> resolveAPIClientPlan cannot infer the refresh op and errors
	// before any file write or openapi-typescript spawn.
	ambiguousDoc := &openapi3.T{Paths: openapi3.NewPaths(
		openapi3.WithPath("/auth/login", &openapi3.PathItem{Post: buildTokenOp("Login", tokenResp, []string{"email"})}),
		openapi3.WithPath("/auth/refresh", &openapi3.PathItem{Post: buildTokenOp("Refresh", tokenResp, []string{"refresh_token"})}),
	)}
	fs := &yongol.Fullstack{
		Manifest: &manifest.ProjectConfig{
			Backend: manifest.Backend{Auth: &manifest.Auth{Mode: "bearer"}},
			Frontend: manifest.Frontend{Auth: &manifest.FrontendAuth{
				TokenField: "access_token", RefreshField: "refresh_token",
			}},
			Domains: map[string]manifest.DomainConfig{
				"public": {OpenAPI: "api/openapi.yaml", Frontend: "frontend", RoutePrefix: "/api"},
			},
		},
		DomainOpenAPIDocs: map[string]*openapi3.T{"public": ambiguousDoc},
	}

	err := generateDomainFrontends(fs, t.TempDir())
	if err == nil {
		t.Fatal("ambiguous refresh-op in a domain: want propagated generate error, got nil")
	}
	if !strings.Contains(err.Error(), "refresh_op") {
		t.Errorf("error should advise declaring frontend.auth.refresh_op, got: %v", err)
	}
}
