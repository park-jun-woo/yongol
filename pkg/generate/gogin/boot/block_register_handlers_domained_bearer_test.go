//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what blockRegisterHandlersDomained(bearer) — 도메인별 publicOps + BearerAuthStrict + middleware import 검증

package boot

import (
	"strings"
	"testing"
)

func TestBlockRegisterHandlersDomained_Bearer(t *testing.T) {
	block := blockRegisterHandlers(domainedFS([]string{"bearerAuth"}), "example.com/app")
	body := strings.Join(block.Lines, "\n")
	imp := strings.Join(block.Imports, "\n")
	for _, must := range []string{
		"publicPublicOps := map[string]bool{",
		`"Login": true,`,
		"adminPublicOps := map[string]bool{",
		"middleware.BearerAuthStrictPublic(publicPublicOps)",
		"middleware.BearerAuthStrictAdmin(adminPublicOps)",
	} {
		if !strings.Contains(body, must) {
			t.Errorf("missing %q in:\n%s", must, body)
		}
	}
	if !strings.Contains(imp, "/internal/middleware") {
		t.Errorf("bearer → must import middleware:\n%s", imp)
	}
}
