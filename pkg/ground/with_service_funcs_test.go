//ff:func feature=rule type=test-helper control=sequence
//ff:what withServiceFuncs — SSaC ServiceFunc 슬라이스를 Fullstack.ServiceFuncs 에 append 하는 option

package ground

import (
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// withServiceFuncs attaches parsed SSaC service funcs.
func withServiceFuncs(funcs ...ssac.ServiceFunc) func(*yongol.Fullstack) {
	return func(fs *yongol.Fullstack) { fs.ServiceFuncs = append(fs.ServiceFuncs, funcs...) }
}
