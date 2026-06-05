//ff:func feature=stml-parse type=util control=selection dimension=1
//ff:what 현재 토큰을 소비하고 반환한다 (EOF에서는 위치를 더 진행하지 않음)
package stml

// advance consumes and returns the current token.
func (p *guardParser) advance() guardToken {
	t := p.toks[p.pos]
	if p.pos < len(p.toks)-1 {
		p.pos++
	}
	return t
}
