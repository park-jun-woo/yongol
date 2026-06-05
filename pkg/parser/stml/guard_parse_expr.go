//ff:func feature=stml-parse type=parser control=iteration dimension=1
//ff:what guard := term (("&&"|"||") term)* — 논리 결합을 좌결합으로 파싱한다
package stml

// parseGuardExpr parses the top-level guard rule: term (("&&"|"||") term)*.
// Logical operators are left-associative.
func (p *guardParser) parseGuardExpr() (*GuardExpr, error) {
	left, err := p.parseGuardTerm()
	if err != nil {
		return nil, err
	}
	for p.peek().kind == tokAnd || p.peek().kind == tokOr {
		opTok := p.advance()
		right, err := p.parseGuardTerm()
		if err != nil {
			return nil, err
		}
		left = &GuardExpr{Kind: GuardBinary, Op: opTok.text, Left: left, Right: right}
	}
	return left, nil
}
