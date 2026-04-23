//ff:func feature=gen-react type=accessor control=sequence
//ff:what fsOpenAPIDoc — Fullstack.OpenAPIDoc 접근자 (nil-safe)

package react

import (
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// fsOpenAPIDoc returns the parsed OpenAPI document or nil.
func fsOpenAPIDoc(fs *yongol.Fullstack) *openapi3.T {
	if fs == nil {
		return nil
	}
	return fs.OpenAPIDoc
}
