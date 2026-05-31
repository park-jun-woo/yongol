//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what TestWriteJSONBUnmarshalScaffolding_ZeroCov — 필드별 Unmarshal 블록 생성
package ssac

import (
	"strings"
	"testing"
)

func TestWriteJSONBUnmarshalScaffolding_ZeroCov(t *testing.T) {
	var sb strings.Builder
	writeJSONBUnmarshalScaffolding(&sb, []jsonbFieldAlias{
		{jsonName: "meta", apiField: "Meta", dbField: "Metadata", localVar: "metaLocal"},
	})
	out := sb.String()
	for _, want := range []string{
		"var metaLocal map[string]interface{}",
		"if len(row.Metadata) > 0 {",
		"json.Unmarshal(row.Metadata, &metaLocal)",
		"return nil, err",
	} {
		if !strings.Contains(out, want) {
			t.Errorf("missing %q in:\n%s", want, out)
		}
	}

	// Empty input → empty output.
	var sb2 strings.Builder
	writeJSONBUnmarshalScaffolding(&sb2, nil)
	if sb2.Len() != 0 {
		t.Errorf("expected no output for empty input, got %q", sb2.String())
	}
}
