//ff:func feature=validate type=test control=sequence topic=config-check
//ff:what SEC-04 positive — override key 가 operationId 에 없으면 ERROR

package openapi_manifest

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestSEC04_OverrideKeyMissing(t *testing.T) {
	fs := &yongol.Fullstack{
		OpenAPIDoc: buildDocWithOps([]string{"UploadAvatar"}),
		Manifest: &manifest.ProjectConfig{
			Backend: manifest.Backend{
				HTTP: &manifest.HTTPConfig{
					Overrides: map[string]manifest.HTTPOverride{
						"UploadAvatarTypo": {BodyLimit: "5MiB"},
					},
				},
			},
		},
	}
	diags := sec04HTTPOverridesOperationID(fs)
	if len(diags) != 1 {
		t.Fatalf("expected 1 diag, got %d: %+v", len(diags), diags)
	}
}
