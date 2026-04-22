//ff:func feature=rule type=rule control=sequence
//ff:what IsSensitiveCol — 현재 컬럼이 @sensitive인지 확인
package rule

import "github.com/park-jun-woo/toulmin/pkg/toulmin"

// IsSensitiveCol checks if the current column has @sensitive annotation.
// Ground.Flags["sensitive"] must be set by the caller.
func IsSensitiveCol(ctx toulmin.Context, _ toulmin.Specs) (bool, any) {
	g, _ := ctx.Get("ground")
	ground := g.(*Ground)
	return ground.Flags["sensitive"], nil
}
