//ff:func feature=validate type=util control=iteration dimension=1 topic=ssac-structural
//ff:what s71BuildScope — seqIdx 시점까지의 유효 변수 scope 구축 (request 항상, currentUser @auth 이후, step 결과 변수, message @subscribe)

package ssac

import (
	parsessac "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func s71BuildScope(fn parsessac.ServiceFunc, upto int) map[string]bool {
	scope := map[string]bool{"request": true}
	if fn.Subscribe != nil {
		scope["message"] = true
	}
	for i := 0; i < upto && i < len(fn.Sequences); i++ {
		seq := fn.Sequences[i]
		if seq.Type == parsessac.SeqAuth {
			scope["currentUser"] = true
		}
		if seq.Result != nil && seq.Result.Var != "" {
			scope[seq.Result.Var] = true
		}
	}
	return scope
}
