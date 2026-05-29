//ff:func feature=gen-gogin type=util control=iteration dimension=1
//ff:what requiredSet — schema.Required 슬라이스를 lookup map 으로 평탄화

package ssac

import "github.com/getkin/kin-openapi/openapi3"

// requiredSet flattens schema.Required into a lookup set. Returns an
// empty (non-nil) map when schema is nil or has no required list so
// callers can index without a nil check.
func requiredSet(schema *openapi3.Schema) map[string]bool {
	if schema == nil {
		return map[string]bool{}
	}
	out := make(map[string]bool, len(schema.Required))
	for _, r := range schema.Required {
		out[r] = true
	}
	return out
}
