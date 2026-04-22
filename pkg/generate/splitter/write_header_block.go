//ff:func feature=gen-splitter type=util control=sequence
//ff:what writeHeaderBlock — 원본 generator 배너를 분할 파일 선두에 복제 (빈 header 는 no-op)
package splitter

import (
	"bytes"
	"strings"
)

// writeHeaderBlock appends the generator banner followed by a blank
// separator line to buf. It is a no-op when header is empty so files
// written outside the normal split flow stay compact.
func writeHeaderBlock(buf *bytes.Buffer, header string) {
	if header == "" {
		return
	}
	buf.WriteString(header)
	if !strings.HasSuffix(header, "\n") {
		buf.WriteByte('\n')
	}
	buf.WriteByte('\n')
}
