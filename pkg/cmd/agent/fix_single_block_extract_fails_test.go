//ff:func feature=agent type=test control=sequence
//ff:what TestFixSingleBlock — default 레이어 / extract 실패→generateNewBlock / 블록추출 후 LLM 에러 분기 검증
package agent

import (
	"bytes"
	"strings"
	"testing"
)

func TestFixSingleBlockExtractFails(t *testing.T) {
	// operationId not present → extract error → generateNewBlock, which fails to
	// find the SSaC file and returns false. Exercises the extractErr branch.
	content := "paths: {}\n"
	var out bytes.Buffer
	got := fixSingleBlock(&out, Config{}, layerOpenAPI, "api/openapi.yaml", t.TempDir()+"/api/openapi.yaml",
		&content, "MissingOp", "desc", "path", []string{"err"})
	if got {
		t.Fatal("expected false when block missing and SSaC unavailable")
	}
	if !strings.Contains(out.String(), "skipped generate") {
		t.Errorf("expected generate-skip message, got: %q", out.String())
	}
}
