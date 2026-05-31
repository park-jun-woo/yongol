//ff:func feature=agent type=test control=sequence
//ff:what TestExtractPathBlockForOp — operationId로 path 블록 추출, 미존재/잘못된 YAML 처리 검증
package agent

import (
	"strings"
	"testing"
)

func TestExtractPathBlockForOpNonMapEntries(t *testing.T) {
	// A scalar path value (not a map) hits the methods type-assert continue;
	// a scalar method detail (not a map) hits the detail type-assert continue.
	// The well-formed /users path still resolves CreateUser.
	content := `paths:
  /scalar: "just a string"
  /users:
    summary: "a string detail"
    post:
      operationId: CreateUser`
	block := extractPathBlockForOp(content, "CreateUser")
	if block == "" {
		t.Fatal("expected non-empty block despite non-map siblings")
	}
	if !strings.Contains(block, "operationId: CreateUser") {
		t.Errorf("block = %q, want CreateUser", block)
	}
}
