//ff:func feature=gen-splitter type=test control=sequence
//ff:what splitter 내부 함수 by-name 직접 호출 테스트 — tsma content-aware 귀속용
package splitter

import (
	"bytes"
	"go/ast"
	"strings"
	"testing"
)

func TestWriteImportBlock_ZeroCov(t *testing.T) {
	var buf bytes.Buffer
	imps := []*ast.ImportSpec{{Path: &ast.BasicLit{Value: `"fmt"`}}}
	writeImportBlock(&buf, imps)
	if !strings.Contains(buf.String(), "import (") {
		t.Errorf("import block wrong: %q", buf.String())
	}
	var empty bytes.Buffer
	writeImportBlock(&empty, nil)
	if empty.Len() != 0 {
		t.Errorf("empty imports should be no-op")
	}
}
