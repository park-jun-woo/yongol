//ff:func feature=rule type=rule control=sequence
//ff:what VarDeclared — 변수가 선언 후 사용되는지 검증
package rule

import "github.com/park-jun-woo/toulmin/pkg/toulmin"

// VarDeclared checks that a variable is declared before use.
// claim: string (variable name). Returns (true, evidence) when NOT declared.
func VarDeclared(ctx toulmin.Context, specs toulmin.Specs) (bool, any) {
	if len(specs) == 0 {
		return false, nil
	}
	s, ok := specs[0].(*VarDeclaredSpec)
	if !ok || s == nil {
		return false, nil
	}
	g, _ := ctx.Get("ground")
	ground := g.(*Ground)
	c, _ := ctx.Get("claim")
	name, _ := c.(string)
	if ground.Vars[name] {
		return false, nil
	}
	return true, &Evidence{Rule: s.Rule, Level: s.Level, Ref: name, Message: s.Message}
}
