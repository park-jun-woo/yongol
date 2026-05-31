//ff:func feature=validate type=util control=iteration dimension=1 topic=ddl-structural
//ff:what scanValidateLineForTerminator — 단일 라인에서 unquoted `;` 여부와 인용 상태 전환 반환

package ddl

// scanValidateLineForTerminator walks ln one character at a time,
// returning (done, nextInSingle). done=true when an unquoted `;` is
// encountered; otherwise the caller carries nextInSingle into the next
// line. Doubled `”` inside a literal counts as an escaped quote.
func scanValidateLineForTerminator(ln string, inSingle bool) (bool, bool) {
	k := 0
	for k < len(ln) {
		ch := ln[k]
		if ch == ';' && !inSingle {
			return true, inSingle
		}
		if ch != '\'' {
			k++
			continue
		}
		if inSingle && k+1 < len(ln) && ln[k+1] == '\'' {
			k += 2
			continue
		}
		inSingle = !inSingle
		k++
	}
	return false, inSingle
}
