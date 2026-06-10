//ff:func feature=gen-react type=util control=iteration dimension=2
//ff:what op2xxResponseProps — operation 의 2xx JSON 응답 스키마 최상위 프로퍼티 set

package react

import (
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

// op2xxResponseProps collects the top-level property names of op's 2xx JSON
// response schemas. Non-2xx codes, non-JSON media types and schema-less
// responses contribute nothing. Mirrors validate's XON-60 traversal
// (pkg/validate/openapi_manifest/op_2xx_property_set.go) so generate-time
// refresh-op inference agrees with what XON-60 verified.
func op2xxResponseProps(op *openapi3.Operation) map[string]bool {
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
