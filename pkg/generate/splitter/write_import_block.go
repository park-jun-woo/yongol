//ff:func feature=gen-splitter type=util control=iteration dimension=1
//ff:what writeImportBlock — 선별된 ImportSpec 들을 Go import 블록 문법으로 buf 에 기록
package splitter

import (
	"bytes"
	"go/ast"
)

// writeImportBlock emits a Go import block. With no imports it writes
// nothing; otherwise it produces a parenthesised multi-line block
// terminated by a blank separator. Aliases and blank/dot forms are
// preserved via renderImportSpec.
func writeImportBlock(buf *bytes.Buffer, imports []*ast.ImportSpec) {
	if len(imports) == 0 {
		return
	}
	buf.WriteString("import (\n")
	for _, imp := range imports {
		buf.WriteByte('\t')
		buf.WriteString(renderImportSpec(imp))
		buf.WriteByte('\n')
	}
	buf.WriteString(")\n\n")
}
