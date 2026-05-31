//ff:func feature=agent type=test control=sequence
//ff:what TestFixSingleBlock — default 레이어 / extract 실패→generateNewBlock / 블록추출 후 LLM 에러 분기 검증
package agent

import (
	"bytes"
	"testing"
)

func TestFixSingleBlockHurlExtract(t *testing.T) {
	// layerHurl routes through extractHurlBlock; a missing op triggers the
	// extract-error → generateNewBlock fallback.
	content := "GET https://x\n"
	var out bytes.Buffer
	if fixSingleBlock(&out, Config{}, layerHurl, "tests/x.hurl", t.TempDir()+"/tests/x.hurl",
		&content, "MissingOp", "d", "p", []string{"err"}) {
		t.Fatal("expected false for hurl with missing op + no SSaC")
	}
}
