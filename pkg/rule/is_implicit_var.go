//ff:func feature=rule type=rule control=sequence
//ff:what IsImplicitVar — 변수가 암묵적으로 선언된 것인지 확인
package rule

import "github.com/park-jun-woo/toulmin/pkg/toulmin"

// IsImplicitVar checks if the claimed variable is implicitly declared.
// Ground.Flags["implicit.<name>"] must be set by the caller.
// claim: string (variable name).
func IsImplicitVar(ctx toulmin.Context, _ toulmin.Specs) (bool, any) {
	g, _ := ctx.Get("ground")
	ground := g.(*Ground)
	c, _ := ctx.Get("claim")
	name, _ := c.(string)
	return ground.Flags["implicit."+name], nil
}
