//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what TestWriteConvertListFunc_ZeroCov — convert<Name>List 본문
package ssac

import (
	"strings"
	"testing"
)

func TestWriteConvertListFunc_ZeroCov(t *testing.T) {
	var sb strings.Builder
	writeConvertListFunc(&sb, "Widget")
	out := sb.String()
	for _, want := range []string{
		"func convertWidgetList(rows []db.Widget) ([]api.Widget, error) {",
		"result := make([]api.Widget, len(rows))",
		"item, err := convertWidget(row)",
		"return result, nil",
	} {
		if !strings.Contains(out, want) {
			t.Errorf("missing %q in:\n%s", want, out)
		}
	}
}
