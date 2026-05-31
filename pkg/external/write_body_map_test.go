//ff:func feature=external type=test control=sequence
//ff:what TestWriteBodyMap/writeReturnWithResult/readBodySnippet — 코드생성·스니펫 검증
package external

import (
	"bytes"
	"testing"
)

func TestWriteBodyMap(t *testing.T) {
	var buf bytes.Buffer
	writeBodyMap(&buf, []paramInfo{
		{Name: "name", In: "body"},
		{Name: "age", In: "body"},
	})
	got := buf.String()
	want := "\tbody := map[string]any{\"name\": name, \"age\": age}\n"
	if got != want {
		t.Errorf("writeBodyMap = %q, want %q", got, want)
	}
}
