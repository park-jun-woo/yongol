//ff:type feature=stml-parse type=model
//ff:what guardToken — 단일 어휘 토큰 (종류 + 원문 텍스트)
package stml

// guardToken is a single lexical token with its source text.
type guardToken struct {
	kind guardTokKind
	text string
}
