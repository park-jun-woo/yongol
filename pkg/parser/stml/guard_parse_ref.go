//ff:func feature=stml-parse type=parser control=sequence dimension=1
//ff:what ref := <model>"."<Field> 를 파싱해 GuardRef 를 반환한다
package stml

import "fmt"

// parseGuardRef parses ref := <model> "." <Field> (e.g. "workflow.status").
func (p *guardParser) parseGuardRef() (GuardRef, error) {
	if p.peek().kind != tokIdent {
		return GuardRef{}, fmt.Errorf("expected model reference in guard expression, got %q", p.peek().text)
	}
	model := p.advance().text
	if p.peek().kind != tokDot {
		return GuardRef{}, fmt.Errorf("expected %q after model %q in guard reference", ".", model)
	}
	p.advance() // consume "."
	if p.peek().kind != tokIdent {
		return GuardRef{}, fmt.Errorf("expected field name after %q in guard reference", model+".")
	}
	field := p.advance().text
	return GuardRef{Model: model, Field: field}, nil
}
