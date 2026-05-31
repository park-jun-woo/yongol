//ff:func feature=gen-splitter type=test control=sequence
//ff:what splitter 내부 함수 by-name 직접 호출 테스트 — tsma content-aware 귀속용
package splitter

import (
	"bytes"
	"strings"
	"testing"
)

func TestWriteAnnotationBlock_ZeroCov(t *testing.T) {
	var buf bytes.Buffer
	writeAnnotationBlock(&buf, []string{"ff:func feature=x", "ff:what y"})
	if !strings.Contains(buf.String(), "//ff:func") {
		t.Errorf("annotation block wrong: %q", buf.String())
	}
}
