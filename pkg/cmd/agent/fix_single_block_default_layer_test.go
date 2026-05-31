//ff:func feature=agent type=test control=sequence
//ff:what TestFixSingleBlock — default 레이어 / extract 실패→generateNewBlock / 블록추출 후 LLM 에러 분기 검증
package agent

import (
	"bytes"
	"testing"
)

func TestFixSingleBlockDefaultLayer(t *testing.T) {
	// An unsupported layer returns false before any extraction.
	content := "x"
	var out bytes.Buffer
	if fixSingleBlock(&out, Config{}, layerDDL, "db/x.sql", "/tmp/db/x.sql", &content, "Op", "", "", nil) {
		t.Fatal("expected false for non-split layer")
	}
}
