//ff:func feature=external type=test control=iteration dimension=1
//ff:what TestWriteBodyMap/writeReturnWithResult/readBodySnippet — 코드생성·스니펫 검증
package external

import (
	"bytes"
	"strings"
	"testing"
)

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
