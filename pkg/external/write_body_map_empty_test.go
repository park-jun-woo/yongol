//ff:func feature=external type=test control=sequence
//ff:what TestWriteBodyMap/writeReturnWithResult/readBodySnippet — 코드생성·스니펫 검증
package external

import (
	"bytes"
	"testing"
)

func TestWriteBodyMapEmpty(t *testing.T) {
	var buf bytes.Buffer
	writeBodyMap(&buf, nil)
	want := "\tbody := map[string]any{}\n"
	if got := buf.String(); got != want {
		t.Errorf("writeBodyMap(nil) = %q, want %q", got, want)
	}
}
