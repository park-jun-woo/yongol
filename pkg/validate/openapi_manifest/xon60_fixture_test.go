//ff:func feature=validate type=test-helper control=sequence topic=config-check
//ff:what 테스트 헬퍼 — frontend.auth + operationId→응답 필드 매핑으로 Fullstack 픽스처 생성

package openapi_manifest

import (
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xon60Fixture builds a Fullstack with the given frontend.auth block and an
// OpenAPI doc whose ops each declare the given 2xx response field names.
// Used by the XON-60 unit tests.
func xon60Fixture(auth *pmanifest.FrontendAuth, ops map[string][]string) *yongol.Fullstack {
	return &yongol.Fullstack{
		Manifest: &pmanifest.ProjectConfig{
			Frontend: pmanifest.Frontend{Auth: auth},
		},
		OpenAPIDoc: buildDocWithResponseFields(ops),
	}
}
