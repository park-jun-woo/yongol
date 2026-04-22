//ff:func feature=gen-gogin type=util control=sequence
//ff:what methodGen.assignOp — err 선언용 := / = 연산자 선택 (FirstErr 추적)

package ssac

// assignOp returns := or = for the next error declaration.
// hasNewVar: true when the LHS introduces a new result variable (not _).
func (g *methodGen) assignOp(hasNewVar bool) string {
	if hasNewVar {
		// new variable on LHS → always :=
		g.FirstErr = false
		return ":="
	}
	// err-only — first time :=, then =
	if g.FirstErr {
		g.FirstErr = false
		return ":="
	}
	return "="
}
