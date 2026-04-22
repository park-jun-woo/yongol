//ff:func feature=validate type=util control=iteration dimension=1 topic=ssac-openapi
//ff:what hasResponseSequence — ServiceFunc에 @response 시퀀스 존재 여부

package openapi_ssac

import "github.com/park-jun-woo/yongol/pkg/parser/ssac"

// hasResponseSequence reports whether fn declares a @response sequence.
func hasResponseSequence(fn ssac.ServiceFunc) bool {
	for _, seq := range fn.Sequences {
		if seq.Type == "response" {
			return true
		}
	}
	return false
}
