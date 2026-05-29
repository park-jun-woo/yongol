//ff:func feature=gen-gogin type=test control=sequence topic=response
//ff:what TestBuildFieldResponse_ManifestRef — manifest.* 참조가 Go 리터럴로 변환되는지 검증

package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

// TestBuildFieldResponse_ManifestRef verifies that a manifest.* reference
// in @response fields is resolved to its Go literal (duration → seconds).
func TestBuildFieldResponse_ManifestRef(t *testing.T) {
	g := &methodGen{
		FuncName:      "Login",
		SuccessStatus: 200,
		RespFields: map[string]responseField{
			"access_token": {JSONName: "access_token", GoName: "AccessToken"},
			"expires_in":   {JSONName: "expires_in", GoName: "ExpiresIn"},
		},
		Manifest: &manifest.ProjectConfig{
			Backend: manifest.Backend{
				Auth: &manifest.Auth{AccessTokenTTL: "15m"},
			},
		},
	}
	fields := map[string]string{
		"access_token": "token.AccessToken",
		"expires_in":   "manifest.auth.accessTokenTTL",
	}
	lines, _ := g.buildFieldResponse(fields)
	body := strings.Join(lines, "\n")

	if !strings.Contains(body, "ptrOf(int64(900))") {
		t.Fatalf("manifest.auth.accessTokenTTL should resolve to ptrOf(int64(900)), got:\n%s", body)
	}
}
