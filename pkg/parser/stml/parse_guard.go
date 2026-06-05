//ff:func feature=stml-parse type=parser control=sequence dimension=1
//ff:what 가드 조건 문자열을 EBNF(§3.4)로 파싱해 GuardExpr AST를 반환한다
package stml

import "fmt"

// ParseGuard parses a guard condition string into a GuardExpr AST following the
// EBNF in §3.4 (comparison · logical && / || · negation ! · parentheses; no
// function calls, arithmetic, or ternaries). It returns an error for any input
// that violates the grammar — that error is the basis for validation rule
// TM-17.
func ParseGuard(condition string) (*GuardExpr, error) {
	toks, err := lexGuard(condition)
	if err != nil {
		return nil, err
	}
	p := &guardParser{toks: toks}
	expr, err := p.parseGuardExpr()
	if err != nil {
		return nil, err
	}
	if p.peek().kind != tokEOF {
		return nil, fmt.Errorf("unexpected token %q after guard expression", p.peek().text)
	}
	return expr, nil
}
