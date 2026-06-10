//ff:func feature=validate type=util control=iteration dimension=2 topic=config-check
//ff:what op2xxPropertySet — operation 의 2xx JSON 응답 스키마 최상위 프로퍼티 이름 set 수집

package openapi_manifest

import (
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

// op2xxPropertySet collects the top-level property names declared on the
// JSON schemas of op's 2xx responses. Non-2xx codes, non-JSON media types
// and schema-less responses contribute nothing. Used by XON-60 to check
// that frontend.auth field claims exist in some success response.
func op2xxPropertySet(op *openapi3.Operation) map[string]bool {
	props := map[string]bool{}
	if op == nil || op.Responses == nil {
		return props
	}
	for code, r := range op.Responses.Map() {
		if !strings.HasPrefix(code, "2") || r == nil || r.Value == nil {
			continue
		}
		for ct, mt := range r.Value.Content {
			if !strings.Contains(strings.ToLower(ct), "json") {
				continue
			}
			if mt == nil || mt.Schema == nil || mt.Schema.Value == nil {
				continue
			}
			for name := range mt.Schema.Value.Properties {
				props[name] = true
			}
		}
	}
	return props
}
