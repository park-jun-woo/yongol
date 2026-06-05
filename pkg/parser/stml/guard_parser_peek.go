//ff:func feature=stml-parse type=util control=sequence dimension=1
//ff:what 현재 토큰을 소비하지 않고 반환한다
package stml

// peek returns the current token without consuming it.
func (p *guardParser) peek() guardToken {
	return p.toks[p.pos]
}
