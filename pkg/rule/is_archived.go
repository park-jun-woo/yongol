//ff:func feature=rule type=rule control=sequence
//ff:what IsArchived — 현재 테이블이 @archived인지 확인
package rule

import "github.com/park-jun-woo/toulmin/pkg/toulmin"

// IsArchived checks if the current table is marked @archived.
// Ground.Flags["archived"] must be set by the caller.
func IsArchived(ctx toulmin.Context, _ toulmin.Specs) (bool, any) {
	g, _ := ctx.Get("ground")
	ground := g.(*Ground)
	return ground.Flags["archived"], nil
}
