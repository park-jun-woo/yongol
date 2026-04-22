//ff:func feature=validate type=util control=iteration dimension=1 topic=ssac-openapi
//ff:what extractResponseFieldKeys — ServiceFunc의 명시적 @response 필드 키 반환 (shorthand면 nil)

package openapi_ssac

import "github.com/park-jun-woo/yongol/pkg/parser/ssac"

// extractResponseFieldKeys returns the explicit @response field keys for fn,
// or nil if fn uses shorthand (@response varName) or has no @response.
func extractResponseFieldKeys(fn ssac.ServiceFunc) []string {
	for _, seq := range fn.Sequences {
		if seq.Type != "response" {
			continue
		}
		if seq.Target != "" || len(seq.Fields) == 0 {
			return nil
		}
		keys := make([]string, 0, len(seq.Fields))
		for k := range seq.Fields {
			keys = append(keys, k)
		}
		return keys
	}
	return nil
}
