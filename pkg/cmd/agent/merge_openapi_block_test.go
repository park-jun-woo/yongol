//ff:func feature=agent type=test control=sequence
//ff:what TestMergeOpenAPIBlock — 잘못된 YAML/operationId 누락 거부 및 정상 머지 검증

package agent

import (
	"strings"
	"testing"
)

func TestMergeOpenAPIBlock(t *testing.T) {
	original := "paths:\n  /old:\n    get:\n      operationId: Old"
	fixed := "/new:\n  get:\n    operationId: New"

	got, err := mergeOpenAPIBlock(original, 1, 4, fixed)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(got, "operationId: New") {
		t.Errorf("merged missing new op: %q", got)
	}

	if _, err := mergeOpenAPIBlock(original, 1, 4, "::: not yaml :::"); err == nil {
		t.Error("expected error for invalid YAML")
	}
	if _, err := mergeOpenAPIBlock(original, 1, 4, "get:\n  summary: x"); err == nil {
		t.Error("expected error for missing operationId")
	}
}
