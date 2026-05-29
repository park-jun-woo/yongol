//ff:func feature=agent type=test control=iteration dimension=1
//ff:what TestInsertOpenAPIBlock — paths 섹션 끝에 신규 path 블록 삽입 및 검증 거부 케이스

package agent

import (
	"strings"
	"testing"
)

func TestInsertOpenAPIBlock(t *testing.T) {
	original := "openapi: 3.0.0\npaths:\n  /old:\n    get:\n      operationId: Old\ncomponents:\n  schemas: {}"
	newBlock := "/new:\n  get:\n    operationId: New"

	got, err := insertOpenAPIBlock(original, newBlock)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(got, "operationId: New") {
		t.Errorf("result missing new block: %q", got)
	}
	// New block inserted before the top-level components section.
	if strings.Index(got, "operationId: New") > strings.Index(got, "components:") {
		t.Errorf("new block should be inside paths before components: %q", got)
	}

	if _, err := insertOpenAPIBlock(original, ":::bad yaml"); err == nil {
		t.Error("expected error for invalid YAML")
	}
	if _, err := insertOpenAPIBlock(original, "/new:\n  get:\n    summary: x"); err == nil {
		t.Error("expected error for missing operationId")
	}
	if _, err := insertOpenAPIBlock("openapi: 3.0.0\ninfo: {}", newBlock); err == nil {
		t.Error("expected error when 'paths:' not found")
	}
}
