//ff:func feature=validate type=test control=sequence topic=config-check
//ff:what SEC-04 테스트 — overrides.<key> 가 operationId 에 존재하는지 교차 검증

package openapi_manifest

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"

	"github.com/park-jun-woo/yongol/pkg/yongol"
	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

func buildDocWithOps(opIDs []string) *openapi3.T {
	doc := &openapi3.T{
		OpenAPI: "3.0.0",
		Info:    &openapi3.Info{Title: "t", Version: "1"},
		Paths:   openapi3.NewPaths(),
	}
	for i, id := range opIDs {
		path := "/ep" + string(rune('a'+i))
		pi := &openapi3.PathItem{
			Post: &openapi3.Operation{OperationID: id, Responses: openapi3.NewResponses()},
		}
		doc.Paths.Set(path, pi)
	}
	return doc
}

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

func TestSEC04_OverrideKeyExists(t *testing.T) {
	fs := &yongol.Fullstack{
		OpenAPIDoc: buildDocWithOps([]string{"UploadAvatar"}),
		Manifest: &manifest.ProjectConfig{
			Backend: manifest.Backend{
				HTTP: &manifest.HTTPConfig{
					Overrides: map[string]manifest.HTTPOverride{
						"UploadAvatar": {BodyLimit: "5MiB"},
					},
				},
			},
		},
	}
	if diags := sec04HTTPOverridesOperationID(fs); len(diags) != 0 {
		t.Fatalf("expected 0 diags, got %d: %+v", len(diags), diags)
	}
}

func TestSEC04_NoHTTPConfig(t *testing.T) {
	fs := &yongol.Fullstack{
		OpenAPIDoc: buildDocWithOps([]string{"Any"}),
		Manifest:   &manifest.ProjectConfig{},
	}
	if diags := sec04HTTPOverridesOperationID(fs); diags != nil {
		t.Fatalf("expected nil diags, got %+v", diags)
	}
}
