//ff:func feature=rule type=rule control=sequence
//ff:what IsCustomTS — 필드가 custom.ts에 존재하는지 확인
package rule

import "github.com/park-jun-woo/toulmin/pkg/toulmin"

// IsCustomTS checks if the claimed field exists in custom.ts.
// Ground.Flags["customTS.<name>"] must be set by the caller.
// claim: string (field name).
func IsCustomTS(ctx toulmin.Context, _ toulmin.Specs) (bool, any) {
	g, _ := ctx.Get("ground")
	ground := g.(*Ground)
	c, _ := ctx.Get("claim")
	name, _ := c.(string)
	return ground.Flags["customTS."+name], nil
}
