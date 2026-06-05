//ff:func feature=stml-parse type=parser control=selection dimension=1
//ff:what atom := ref op value | ref "." lifecycle | "(" guard ")" 를 파싱한다
package stml

import "fmt"

// parseGuardAtom parses atom := ref op value | ref "." lifecycle | "(" guard ")".
func (p *guardParser) parseGuardAtom() (*GuardExpr, error) {
	if p.peek().kind == tokLParen {
		return p.parseGuardGroup()
	}
	ref, err := p.parseGuardRef()
	if err != nil {
		return nil, err
	}
	switch p.peek().kind {
	case tokDot:
		return p.parseGuardLifecycle(ref)
	case tokOp:
		return p.parseGuardCompare(ref)
	default:
		return nil, fmt.Errorf("expected comparison operator or lifecycle after %q in guard expression", ref.Path())
	}
}
