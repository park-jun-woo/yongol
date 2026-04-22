//ff:func feature=rule type=rule control=sequence
//ff:what IsSubscribe — 현재 함수가 @subscribe 핸들러인지 확인
package rule

import "github.com/park-jun-woo/toulmin/pkg/toulmin"

// IsSubscribe checks if the current function is a @subscribe handler.
// Ground.Flags["subscribe"] must be set by the caller.
func IsSubscribe(ctx toulmin.Context, _ toulmin.Specs) (bool, any) {
	g, _ := ctx.Get("ground")
	ground := g.(*Ground)
	return ground.Flags["subscribe"], nil
}
