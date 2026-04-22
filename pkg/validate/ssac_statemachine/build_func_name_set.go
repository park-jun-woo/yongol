//ff:func feature=validate type=util control=iteration dimension=1 topic=states
//ff:what ServiceFunc 리스트에서 이름 set 생성

package ssac_statemachine

import (
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

// buildFuncNameSet returns a set of function names.
func buildFuncNameSet(funcs []ssac.ServiceFunc) map[string]bool {
	m := make(map[string]bool, len(funcs))
	for _, fn := range funcs {
		m[fn.Name] = true
	}
	return m
}
