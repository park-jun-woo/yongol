//ff:func feature=validate type=util control=iteration dimension=1 topic=hurl-openapi
//ff:what jsonSchemaFromResponse — response 의 첫 JSON 미디어 타입 schema 추출

package hurl_openapi

import (
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

// jsonSchemaFromResponse extracts the first JSON-ish media type's
// schema from a response ref, treating absent values as nil.
func jsonSchemaFromResponse(r *openapi3.ResponseRef) *openapi3.Schema {
	if r == nil || r.Value == nil {
		return nil
	}
	for ct, mt := range r.Value.Content {
		if !strings.Contains(strings.ToLower(ct), "json") {
			continue
		}
		if mt == nil || mt.Schema == nil {
			continue
		}
		return mt.Schema.Value
	}
	return nil
}
