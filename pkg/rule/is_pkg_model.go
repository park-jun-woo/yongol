//ff:func feature=rule type=rule control=sequence
//ff:what IsPkgModel — 현재 모델이 pkg 모델인지 확인
package rule

import "github.com/park-jun-woo/toulmin/pkg/toulmin"

// IsPkgModel checks if the current model is a pkg model (no DDL table).
// Ground.Flags["pkgModel"] must be set by the caller.
func IsPkgModel(ctx toulmin.Context, _ toulmin.Specs) (bool, any) {
	g, _ := ctx.Get("ground")
	ground := g.(*Ground)
	return ground.Flags["pkgModel"], nil
}
