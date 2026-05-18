//ff:func feature=gen-gogin type=util control=sequence
//ff:what methodGen.assignOp — err 선언용 := / = 연산자 선택 (FirstErr + DeclaredVars 추적)

package ssac

// assignOp returns := or = for the next error declaration.
// hasNewVar: true when the LHS introduces a new result variable (not _).
// resultVar: the name of the result variable being bound ("" or "_" when
// the call has no result binding). When the same resultVar has already
// been declared in a previous sequence, = is returned to avoid a
// "no new variables on left side of :=" compile error (BUG-069).
func (g *methodGen) assignOp(hasNewVar bool, resultVar string) string {
	if resultVar != "" && resultVar != "_" {
		if g.DeclaredVars[resultVar] {
			// Variable already declared — reuse with plain assignment.
			return "="
		}
		g.DeclaredVars[resultVar] = true
	}
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
