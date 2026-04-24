//ff:func feature=manifest type=parser control=sequence
//ff:what rebuildBufferWithTerminator — INSERT 종결 `;` 까지만 남기고 buffer 재구성
package ddl

import (
	"strings"
)

// rebuildBufferWithTerminator truncates the collected INSERT buffer so it
// ends at the first unquoted `;` on line j (column k). Called only when
// findUnquotedSemicolon reports a terminator on the current line.
// i == j  → the terminator is on the INSERT line itself; reset and take
// lines[i][:k+1]. Otherwise, concatenate lines[i..j-1] verbatim and
// append lines[j][:k+1].
func rebuildBufferWithTerminator(buf *strings.Builder, lines []string, i, j int, ln string, k int) {
	if j == i {
		buf.Reset()
		buf.WriteString(ln[:k+1])
		return
	}
	prev := strings.Join(lines[i:j], "\n")
	buf.Reset()
	buf.WriteString(prev)
	buf.WriteByte('\n')
	buf.WriteString(ln[:k+1])
}
