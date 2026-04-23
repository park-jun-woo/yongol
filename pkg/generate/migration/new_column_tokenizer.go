//ff:func feature=migration type=util control=sequence
//ff:what newColumnTokenizer — 빈 columnTokenizer 반환
package migration

func newColumnTokenizer() *columnTokenizer { return &columnTokenizer{} }
