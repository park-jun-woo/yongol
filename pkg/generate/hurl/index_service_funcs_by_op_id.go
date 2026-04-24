//ff:func feature=gen-hurl type=util control=iteration dimension=1
//ff:what indexServiceFuncsByOpID — SSaC ServiceFunc 를 operationId 키로 맵핑

package hurl

import (
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

// indexServiceFuncsByOpID builds a lookup from operationId (ServiceFunc.Name)
// to the parsed SSaC ServiceFunc. yongol's convention pins the .ssac
// function name to the OpenAPI operationId (e.g. `func Signup() {}` ↔
// operationId: Signup), so direct name match is safe.
func indexServiceFuncsByOpID(funcs []ssac.ServiceFunc) map[string]*ssac.ServiceFunc {
	out := make(map[string]*ssac.ServiceFunc, len(funcs))
	for i := range funcs {
		fn := &funcs[i]
		if fn.Name == "" {
			continue
		}
		out[fn.Name] = fn
	}
	return out
}
