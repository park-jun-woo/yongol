//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what TestWriteConvertFunc_ZeroCov — convert<Name> 본문 (required/optional 분기)
package ssac

import (
	"strings"
	"testing"
)

func TestWriteConvertFunc_ZeroCov(t *testing.T) {
	var sb strings.Builder
	writeConvertFunc(&sb, "Widget", convertSchemaZeroCov(), nil)
	out := sb.String()
	for _, want := range []string{
		"func convertWidget(row db.Widget) (*api.Widget, error) {",
		"return &api.Widget{",
		"Id:",
		"Name:",
		"}, nil",
	} {
		if !strings.Contains(out, want) {
			t.Errorf("missing %q in:\n%s", want, out)
		}
	}
}
