//ff:func feature=validate type=util control=iteration dimension=1 topic=ssac-openapi
//ff:what shorthandResponseTarget — ServiceFunc의 shorthand @response 변수명 반환

package openapi_ssac

import "github.com/park-jun-woo/yongol/pkg/parser/ssac"

// shorthandResponseTarget returns the variable name of a shorthand @response
// sequence (e.g. "@response course" → "course"), or "" when fn has no
// shorthand @response.
func shorthandResponseTarget(fn ssac.ServiceFunc) string {
	for _, seq := range fn.Sequences {
		if seq.Type != "response" {
			continue
		}
		if seq.Target != "" {
			return seq.Target
		}
	}
	return ""
}
