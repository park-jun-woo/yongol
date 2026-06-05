//ff:func feature=validate type=test-helper control=sequence topic=stml-openapi
//ff:what makeFS — 테스트용 Fullstack fixture 생성

package stml_openapi

import (
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// makeFS builds a Fullstack with the given pages and OpenAPI doc. The manifest
// defaults to a frontend-ON config (Lang set) so XMO-10/11/12 coverage rules,
// which gate on frontendEnabled, run under the default fixture.
func makeFS(pages []stml.PageSpec, doc *openapi3.T) *yongol.Fullstack {
	return &yongol.Fullstack{
		SpecsDir:   "/tmp/test-project",
		STMLPages:  pages,
		OpenAPIDoc: doc,
		Manifest:   &manifest.ProjectConfig{Frontend: manifest.Frontend{Lang: "typescript"}},
	}
}
