//ff:func feature=validate type=util control=iteration dimension=1 topic=openapi-ddl
//ff:what responseSeqOf — ServiceFunc 의 @response 시퀀스 반환 (없으면 nil)

package openapi_ddl

import "github.com/park-jun-woo/yongol/pkg/parser/ssac"

// responseSeqOf returns the first @response sequence of fn, or nil when the
// function declares none.
func responseSeqOf(fn *ssac.ServiceFunc) *ssac.Sequence {
	for i := range fn.Sequences {
		if fn.Sequences[i].Type == "response" {
			return &fn.Sequences[i]
		}
	}
	return nil
}
