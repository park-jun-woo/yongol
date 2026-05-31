//ff:func feature=manifest type=parser control=iteration dimension=1
//ff:what findUnquotedSemicolon — 라인 내 단일 인용 밖의 `;` 위치 탐색
package ddl

// findUnquotedSemicolon scans ln character-by-character starting with the
// quote state inSingle and returns the index of the first `;` that falls
// outside a single-quoted literal. Returns (-1, false) when the line
// terminates without an unquoted `;`. Handles doubled `”` as an escaped
// quote inside a literal.
func findUnquotedSemicolon(ln string, inSingle bool) (int, bool) {
	k := 0
	for k < len(ln) {
		ch := ln[k]
		if ch == ';' && !inSingle {
			return k, true
		}
		if ch != '\'' {
			k++
			continue
		}
		// doubled '' inside a literal is an escaped quote
		if inSingle && k+1 < len(ln) && ln[k+1] == '\'' {
			k += 2
			continue
		}
		inSingle = !inSingle
		k++
	}
	return -1, false
}
