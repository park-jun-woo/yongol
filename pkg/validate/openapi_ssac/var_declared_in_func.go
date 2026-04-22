//ff:func feature=validate type=util control=iteration dimension=1 topic=ssac-openapi
//ff:what varDeclaredInFunc — ServiceFunc 시퀀스에서 varName이 Result로 선언되었는지 확인

package openapi_ssac

import "github.com/park-jun-woo/yongol/pkg/parser/ssac"

// varDeclaredInFunc reports whether `varName` is bound as the Result.Var of
// any sequence (get/post/put/delete/call) in fn. Returns false when varName is
// used as a shorthand @response target but never declared upstream.
func varDeclaredInFunc(fn ssac.ServiceFunc, varName string) bool {
	if varName == "" {
		return false
	}
	for _, seq := range fn.Sequences {
		if seq.Result != nil && seq.Result.Var == varName {
			return true
		}
	}
	return false
}
