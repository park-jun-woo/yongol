//ff:type feature=stml-parse type=model
//ff:what guardParser — 가드 재귀하강 파서 상태 (토큰 슬라이스 + 현재 위치)
package stml

// guardParser holds tokens and the current read position for recursive-descent
// parsing of a guard expression.
type guardParser struct {
	toks []guardToken
	pos  int
}
