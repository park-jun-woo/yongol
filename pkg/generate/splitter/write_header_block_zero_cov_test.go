//ff:func feature=gen-splitter type=test control=sequence
//ff:what splitter 내부 함수 by-name 직접 호출 테스트 — tsma content-aware 귀속용
package splitter

import (
	"bytes"
	"strings"
	"testing"
)

func TestWriteHeaderBlock_ZeroCov(t *testing.T) {
	var buf bytes.Buffer
	writeHeaderBlock(&buf, "// Code generated")
	if !strings.Contains(buf.String(), "Code generated") {
		t.Errorf("header missing: %q", buf.String())
	}
	var empty bytes.Buffer
	writeHeaderBlock(&empty, "")
	if empty.Len() != 0 {
		t.Errorf("empty header should be no-op")
	}
}
