//ff:func feature=manifest type=parser control=sequence
//ff:what advanceSingleQuote — 현재 위치가 '' 이스케이프인지 여부로 inSQ 플래그와 다음 인덱스 결정

package ddl

// advanceSingleQuote inspects the character at position i of line (known to
// be a single quote). Returns the updated inSQ flag and the next scan
// position. `''` inside a string literal is a SQL escape and keeps inSQ true
// while skipping the following quote.
func advanceSingleQuote(line string, i int, inSQ bool) (bool, int) {
	if inSQ && i+1 < len(line) && line[i+1] == '\'' {
		return inSQ, i + 2
	}
	return !inSQ, i + 1
}
