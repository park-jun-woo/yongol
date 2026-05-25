//ff:func feature=gen-gogin type=test control=sequence topic=response
//ff:what TestBuildFieldResponse_ManifestRefRefreshToken — refreshTokenTTL 168h → 604800 변환 검증

package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

// TestBuildFieldResponse_ManifestRefRefreshToken verifies refresh token TTL.
func TestBuildFieldResponse_ManifestRefRefreshToken(t *testing.T) {
	g := &methodGen{
		FuncName:      "Login",
		SuccessStatus: 200,
		RespFields: map[string]responseField{
			"refresh_expires_in": {JSONName: "refresh_expires_in", GoName: "RefreshExpiresIn"},
		},
		Manifest: &manifest.ProjectConfig{
			Backend: manifest.Backend{
				Auth: &manifest.Auth{RefreshTokenTTL: "168h"},
			},
		},
	}
	fields := map[string]string{
		"refresh_expires_in": "manifest.auth.refreshTokenTTL",
	}
	lines, _ := g.buildFieldResponse(fields)
	body := strings.Join(lines, "\n")

	if !strings.Contains(body, "ptrOf(int64(604800))") {
		t.Fatalf("manifest.auth.refreshTokenTTL should resolve to ptrOf(int64(604800)), got:\n%s", body)
	}
}
