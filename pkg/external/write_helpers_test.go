//ff:func feature=external type=test control=sequence
//ff:what TestWriteBodyMap/writeReturnWithResult/readBodySnippet — 코드생성·스니펫 검증

package external

import (
	"bytes"
	"strings"
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

func TestWriteBodyMapEmpty(t *testing.T) {
	var buf bytes.Buffer
	writeBodyMap(&buf, nil)
	want := "\tbody := map[string]any{}\n"
	if got := buf.String(); got != want {
		t.Errorf("writeBodyMap(nil) = %q, want %q", got, want)
	}
}

func TestWriteReturnWithResult(t *testing.T) {
	var buf bytes.Buffer
	writeReturnWithResult(&buf, "GET", `"/items/"+itemID`, "nil", "GetItemResponse")
	got := buf.String()

	for _, want := range []string{
		"\tvar resp GetItemResponse\n",
		`c.do(ctx, "GET", "/items/"+itemID, nil, &resp)`,
		"\t\treturn nil, err\n",
		"\treturn &resp, nil\n",
	} {
		if !strings.Contains(got, want) {
			t.Errorf("writeReturnWithResult output missing %q\n got:\n%s", want, got)
		}
	}
}

func TestReadBodySnippet(t *testing.T) {
	tests := []struct {
		name  string
		in    string
		limit int
		want  string
	}{
		{"short", "hello", 100, "hello"},
		{"trimmed", "  hi  ", 100, "hi"},
		{"newlines collapsed", "a\nb\nc", 100, "a b c"},
		{"truncated", "abcdefgh", 3, "abc..."},
		{"empty", "", 10, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := readBodySnippet(strings.NewReader(tt.in), tt.limit)
			if got != tt.want {
				t.Errorf("readBodySnippet(%q, %d) = %q, want %q", tt.in, tt.limit, got, tt.want)
			}
		})
	}
}
