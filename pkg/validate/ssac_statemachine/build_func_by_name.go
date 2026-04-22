//ff:func feature=validate type=util control=iteration dimension=1 topic=states
//ff:what ServiceFunc 리스트에서 이름 → ServiceFunc 맵 생성

package ssac_statemachine

import (
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

// buildFuncByName returns a map of function name → ServiceFunc.
func buildFuncByName(funcs []ssac.ServiceFunc) map[string]ssac.ServiceFunc {
	m := make(map[string]ssac.ServiceFunc, len(funcs))
	for _, fn := range funcs {
		m[fn.Name] = fn
	}
	return m
}
