//ff:func feature=manifest type=parser control=iteration dimension=1
//ff:what trackQuoteState — 라인 끝까지 스캔하여 단일 인용 상태 전환을 누적
package ddl

// trackQuoteState scans the full line and returns the resulting single-
// quote state after processing it. Used when a line does not contain a
// terminating `;` so the outer collector can carry the state into the
// next line.
func trackQuoteState(ln string, inSingle bool) bool {
	k := 0
	for k < len(ln) {
		ch := ln[k]
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
	return inSingle
}
