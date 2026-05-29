//ff:func feature=agent type=test control=sequence
//ff:what TestMergeHurlBlock — HTTP method 라인 누락 시 에러, 정상 시 라인 교체 검증

package agent

import (
	"strings"
	"testing"
)

func TestMergeHurlBlock(t *testing.T) {
	original := "# Old\nGET http://x/a\nHTTP 200"
	fixed := "# New\nPOST http://x/b\nHTTP 201"

	got, err := mergeHurlBlock(original, 0, 3, fixed)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(got, "POST http://x/b") {
		t.Errorf("merged content missing new block: %q", got)
	}

	if _, err := mergeHurlBlock(original, 0, 3, "# New\nno method here"); err == nil {
		t.Fatal("expected error for block missing HTTP method line")
	}
}
