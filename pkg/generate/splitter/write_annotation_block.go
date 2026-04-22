//ff:func feature=gen-splitter type=util control=iteration dimension=1
//ff:what writeAnnotationBlock — //ff:func / //ff:what 라인들을 buf 에 기록
package splitter

import (
	"bytes"
	"fmt"
)

// writeAnnotationBlock writes one filefunc annotation per line to buf.
// Each entry in annotations is the raw text without the leading "//";
// this function adds the prefix and trailing newline so callers do not
// have to concern themselves with comment syntax.
func writeAnnotationBlock(buf *bytes.Buffer, annotations []string) {
	for _, ann := range annotations {
		fmt.Fprintf(buf, "//%s\n", ann)
	}
}
