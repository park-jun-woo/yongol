//ff:func feature=agent type=test control=sequence
//ff:what TestInsertHurlBlock — HTTP method 누락 거부 및 파일 끝 추가 검증

package agent

import (
	"strings"
	"testing"
)

func TestInsertHurlBlock(t *testing.T) {
	original := "# Existing\nGET http://x/a\nHTTP 200\n"
	newBlock := "# New\nPOST http://x/b\nHTTP 201\n"

	got, err := insertHurlBlock(original, newBlock)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(got, "GET http://x/a") || !strings.Contains(got, "POST http://x/b") {
		t.Errorf("result missing blocks: %q", got)
	}
	if strings.Index(got, "GET") > strings.Index(got, "POST") {
		t.Errorf("new block should be appended after existing: %q", got)
	}
	if !strings.HasSuffix(got, "\n") {
		t.Errorf("result should end with newline: %q", got)
	}

	if _, err := insertHurlBlock(original, "# New\nno method"); err == nil {
		t.Error("expected error for block missing HTTP method line")
	}
}
