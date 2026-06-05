//ff:func feature=stml-parse type=parser control=sequence dimension=1
//ff:what ref "." lifecycle (loading|error|empty) 를 파싱해 GuardLifecycle 노드를 만든다
package stml

import "fmt"

// parseGuardLifecycle parses the lifecycle tail: "." (loading|error|empty)
// following a ref. The caller has verified the current token is ".".
func (p *guardParser) parseGuardLifecycle(ref GuardRef) (*GuardExpr, error) {
	p.advance() // consume "."
	if p.peek().kind != tokIdent {
		return nil, fmt.Errorf("expected lifecycle keyword after %q in guard expression", ref.Path()+".")
	}
	kw := p.advance().text
	if kw != "loading" && kw != "error" && kw != "empty" {
		return nil, fmt.Errorf("invalid lifecycle %q in guard expression (expected loading, error, or empty)", kw)
	}
	return &GuardExpr{Kind: GuardLifecycle, Ref: ref, Lifecycle: kw}, nil
}
