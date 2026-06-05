//ff:func feature=stml-parse type=parser control=sequence dimension=1
//ff:what "(" guard ")" 괄호 그룹을 파싱해 GuardGroup 노드로 감싼다
package stml

import "fmt"

// parseGuardGroup parses "(" guard ")" and wraps the inner expression in a
// GuardGroup node. The caller has verified the current token is "(".
func (p *guardParser) parseGuardGroup() (*GuardExpr, error) {
	p.advance() // consume "("
	inner, err := p.parseGuardExpr()
	if err != nil {
		return nil, err
	}
	if p.peek().kind != tokRParen {
		return nil, fmt.Errorf("expected %q to close group in guard expression", ")")
	}
	p.advance() // consume ")"
	return &GuardExpr{Kind: GuardGroup, Operand: inner}, nil
}
