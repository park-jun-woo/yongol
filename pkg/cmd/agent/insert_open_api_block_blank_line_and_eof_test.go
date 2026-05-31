//ff:func feature=agent type=test control=sequence
//ff:what TestInsertOpenAPIBlock — paths 섹션 끝에 신규 path 블록 삽입 및 검증 거부 케이스
package agent

import (
	"strings"
	"testing"
)

func TestInsertOpenAPIBlockBlankLineAndEOF(t *testing.T) {
	// A blank line inside the paths section exercises the empty-line continue;
	// paths being the last top-level section makes insertAt default to len(lines)
	// so the new block is appended at the end.
	original := "openapi: 3.0.0\npaths:\n  /old:\n    get:\n      operationId: Old\n\n"
	newBlock := "/new:\n  get:\n    operationId: New"
	got, err := insertOpenAPIBlock(original, newBlock)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(got, "operationId: New") {
		t.Errorf("result missing new block: %q", got)
	}
	// New block appended after the existing op (no later top-level section).
	if strings.Index(got, "operationId: New") < strings.Index(got, "operationId: Old") {
		t.Errorf("new block should be appended after Old: %q", got)
	}
}
