//ff:func feature=validate type=util control=iteration dimension=1 topic=ssac-structural
//ff:what varTypes — func 전체에 걸쳐 result 변수명 → raw type 맵 생성

package ssac

import (
	parsessac "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

// varTypes maps result variable names to their (raw) types across the func.
func varTypes(fn parsessac.ServiceFunc) map[string]string {
	out := make(map[string]string)
	for _, seq := range fn.Sequences {
		if seq.Result != nil && seq.Result.Var != "" {
			out[seq.Result.Var] = seq.Result.Type
		}
	}
	return out
}
