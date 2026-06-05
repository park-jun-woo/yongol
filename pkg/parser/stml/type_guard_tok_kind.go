//ff:type feature=stml-parse type=model
//ff:what guardTokKind — 가드 토크나이저가 산출하는 어휘 토큰 종류 enum
package stml

// guardTokKind discriminates lexical token categories of a guard expression.
type guardTokKind int

const (
	tokEOF    guardTokKind = iota
	tokAnd                 // &&
	tokOr                  // ||
	tokNot                 // !
	tokLParen              // (
	tokRParen              // )
	tokDot                 // .
	tokOp                  // = != > < >= <=
	tokIdent               // model / field / lifecycle keyword / enum-literal / number
	tokString              // 'quoted' or "quoted" value
)
