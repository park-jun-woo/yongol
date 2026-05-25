//ff:func feature=gen-gogin type=test control=sequence topic=response
//ff:what TestBuildFieldResponse_ManifestRefNilManifest — nil manifest 시 manifest.* 미해석 검증

package ssac

import (
	"strings"
	"testing"
)

// TestBuildFieldResponse_ManifestRefNilManifest verifies that when the
// manifest is nil, manifest.* references are left unresolved (pass-through).
func TestBuildFieldResponse_ManifestRefNilManifest(t *testing.T) {
	g := &methodGen{
		FuncName:      "Login",
		SuccessStatus: 200,
		RespFields: map[string]responseField{
			"expires_in": {JSONName: "expires_in", GoName: "ExpiresIn"},
		},
		Manifest: nil,
	}
	fields := map[string]string{
		"expires_in": "manifest.auth.accessTokenTTL",
	}
	lines, _ := g.buildFieldResponse(fields)
	body := strings.Join(lines, "\n")

	if strings.Contains(body, "ptrOf") {
		t.Fatalf("nil manifest should not produce ptrOf, got:\n%s", body)
	}
}
