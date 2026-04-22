//ff:func feature=gen-hurl type=util control=iteration dimension=1
//ff:what firstMediaType — Content map에서 첫 번째 MediaType 반환
package hurl

import (
	"github.com/getkin/kin-openapi/openapi3"
)

// firstMediaType returns the first media type from an OpenAPI content map.
func firstMediaType(content openapi3.Content) *openapi3.MediaType {
	for _, v := range content {
		return v
	}
	return nil
}
