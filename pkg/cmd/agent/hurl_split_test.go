//ff:func feature=agent type=test control=sequence
//ff:what TestExtractHurlBlock — "# OperationId" 기준 블록 추출, 경계/미존재 op 처리 검증
package agent

import (
	"strings"
	"testing"
)

func TestExtractHurlBlock(t *testing.T) {
	content := `# Login
POST http://x/auth/login
HTTP 200

# ListUsers
GET http://x/users
HTTP 200
`
	block, start, end, err := extractHurlBlock(content, "Login")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(block, "# Login") || !strings.Contains(block, "POST http://x/auth/login") {
		t.Errorf("block = %q", block)
	}
	if strings.Contains(block, "ListUsers") {
		t.Errorf("block should stop before next comment: %q", block)
	}
	if start != 0 {
		t.Errorf("start = %d, want 0", start)
	}
	// End must exclude the trailing blank line before the next block.
	if end <= start {
		t.Errorf("end = %d, want > start %d", end, start)
	}

	// Unknown operationId errors.
	if _, _, _, err := extractHurlBlock(content, "Nope"); err == nil {
		t.Error("expected error for missing operationId comment")
	}
}
