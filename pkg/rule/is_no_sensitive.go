//ff:func feature=rule type=rule control=sequence
//ff:what IsNoSensitive — 현재 컬럼이 @nosensitive인지 확인
package rule

import "github.com/park-jun-woo/toulmin/pkg/toulmin"

// IsNoSensitive checks if the current column has @nosensitive annotation.
// Ground.Flags["nosensitive"] must be set by the caller.
func IsNoSensitive(ctx toulmin.Context, _ toulmin.Specs) (bool, any) {
	g, _ := ctx.Get("ground")
	ground := g.(*Ground)
	return ground.Flags["nosensitive"], nil
}
