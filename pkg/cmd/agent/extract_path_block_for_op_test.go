//ff:func feature=agent type=test control=iteration dimension=3
//ff:what TestExtractPathBlockForOp — operationId로 path 블록 추출, 미존재/잘못된 YAML 처리 검증

package agent

import (
	"strings"
	"testing"
)

func TestExtractPathBlockForOp(t *testing.T) {
	content := `paths:
  /users:
    get:
      operationId: ListUsers
    post:
      operationId: CreateUser
  /orgs:
    get:
      operationId: ListOrgs`

	block := extractPathBlockForOp(content, "CreateUser")
	if block == "" {
		t.Fatal("expected non-empty block for CreateUser")
	}
	if !strings.Contains(block, "/users") || !strings.Contains(block, "operationId: CreateUser") {
		t.Errorf("block = %q, want /users + CreateUser", block)
	}
	if strings.Contains(block, "ListUsers") {
		t.Errorf("block should isolate the post method only: %q", block)
	}

	if got := extractPathBlockForOp(content, "Nonexistent"); got != "" {
		t.Errorf("unknown op = %q, want empty", got)
	}
	if got := extractPathBlockForOp(":::bad yaml", "X"); got != "" {
		t.Errorf("invalid yaml = %q, want empty", got)
	}
}
