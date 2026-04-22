//ff:func feature=gen-react type=util control=sequence
//ff:what Fullstack → FrontendTheme 조회 (없으면 nil 반환)

package react

import (
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/yongol"
	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

// manifestTheme is a package-local alias so the rest of the react package
// doesn't need a direct manifest import.
type manifestTheme = manifest.FrontendTheme

// resolveTheme extracts the Frontend.Theme pointer (may be nil).
func resolveTheme(fs *yongol.Fullstack) *manifestTheme {
	if fs == nil || fs.Manifest == nil {
		return nil
	}
	return fs.Manifest.Frontend.Theme
}

// findOpenAPISpec returns the path to the OpenAPI yaml used by openapi-typescript.
// Convention: <specsDir>/api/openapi.yaml.
func findOpenAPISpec(fs *yongol.Fullstack) string {
	if fs == nil || fs.SpecsDir == "" {
		return ""
	}
	return fs.SpecsDir + "/api/openapi.yaml"
}

// fsOpenAPIDoc returns the parsed OpenAPI document or nil.
func fsOpenAPIDoc(fs *yongol.Fullstack) *openapi3.T {
	if fs == nil {
		return nil
	}
	return fs.OpenAPIDoc
}
