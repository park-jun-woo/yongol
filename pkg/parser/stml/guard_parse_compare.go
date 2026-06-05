//ff:func feature=stml-parse type=parser control=sequence dimension=1
//ff:what ref op value (= != > < >= <=) 비교식을 파싱해 GuardCompare 노드를 만든다
package stml

import "fmt"

// parseGuardCompare parses the comparison tail: op value following a ref. The
// caller has verified the current token is a comparison operator.
func (p *guardParser) parseGuardCompare(ref GuardRef) (*GuardExpr, error) {
	op := p.advance().text
	v := p.peek()
	if v.kind != tokIdent && v.kind != tokString {
		return nil, fmt.Errorf("expected value after %q in guard comparison, got %q", op, v.text)
	}
	p.advance()
	return &GuardExpr{Kind: GuardCompare, Ref: ref, Op: op, Value: v.text}, nil
}
