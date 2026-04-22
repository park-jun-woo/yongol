//ff:func feature=rule type=rule control=sequence
//ff:what CoverageCheck — 정의된 항목이 사용되고 있는지 검증
package rule

import "github.com/park-jun-woo/toulmin/pkg/toulmin"

// CoverageCheck checks that a defined item is used somewhere.
// claim: string (defined item name). Returns (true, evidence) when NOT used.
func CoverageCheck(ctx toulmin.Context, specs toulmin.Specs) (bool, any) {
	if len(specs) == 0 {
		return false, nil
	}
	s, ok := specs[0].(*CoverageCheckSpec)
	if !ok || s == nil {
		return false, nil
	}
	g, _ := ctx.Get("ground")
	ground := g.(*Ground)
	c, _ := ctx.Get("claim")
	name, _ := c.(string)
	if used, ok := ground.Lookup[s.LookupKey]; ok && used[name] {
		return false, nil
	}
	return true, &Evidence{Rule: s.Rule, Level: s.Level, Ref: name, Message: s.Message}
}
