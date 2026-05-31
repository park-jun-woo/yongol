//ff:func feature=agent type=test control=sequence
//ff:what TestFixSingleBlock — default 레이어 / extract 실패→generateNewBlock / 블록추출 후 LLM 에러 분기 검증
package agent

import (
	"bytes"
	"testing"
)

func TestFixSingleBlockRegoExtract(t *testing.T) {
	// layerRego routes through extractRegoBlock; a missing op triggers the
	// extract-error → generateNewBlock fallback.
	content := "package x\n"
	var out bytes.Buffer
	if fixSingleBlock(&out, Config{}, layerRego, "policy/x.rego", t.TempDir()+"/policy/x.rego",
		&content, "MissingOp", "d", "p", []string{"err"}) {
		t.Fatal("expected false for rego with missing op + no SSaC")
	}
}
