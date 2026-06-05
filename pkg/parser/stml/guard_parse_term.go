//ff:func feature=stml-parse type=parser control=selection dimension=1
//ff:what term := "!"? atom — 선행 부정(!)을 단항 노드로 감싼다
package stml

// parseGuardTerm parses term := "!"? atom. A leading "!" wraps the atom in a
// unary negation node.
func (p *guardParser) parseGuardTerm() (*GuardExpr, error) {
	if p.peek().kind == tokNot {
		p.advance()
		operand, err := p.parseGuardAtom()
		if err != nil {
			return nil, err
		}
		return &GuardExpr{Kind: GuardUnary, Op: "!", Operand: operand}, nil
	}
	return p.parseGuardAtom()
}
