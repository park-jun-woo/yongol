//ff:func feature=migration type=util control=sequence
//ff:what crossCategoryCast — 숫자 카테고리 ↔ 텍스트 카테고리 변환 판정
package migration

// crossCategoryCast returns true when fromBase and toBase belong to
// different broad categories (numeric vs text).
func crossCategoryCast(fromBase, toBase string) bool {
	num := map[string]bool{"INTEGER": true, "BIGINT": true, "SMALLINT": true, "NUMERIC": true, "REAL": true, "DOUBLE PRECISION": true}
	text := map[string]bool{"TEXT": true, "VARCHAR": true, "CHAR": true}
	if num[fromBase] && text[toBase] {
		return true
	}
	if text[fromBase] && num[toBase] {
		return true
	}
	return false
}
