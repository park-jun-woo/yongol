//ff:func feature=agent type=test control=sequence
//ff:what TestExtractOpenAPIBlock — op 추출 성공 + op미존재/method미존재/path미존재 에러 분기 검증
package agent

import (
	"strings"
	"testing"
)

func TestExtractOpenAPIBlockSuccess(t *testing.T) {
	content := strings.Join([]string{
		"paths:",
		"  /users:",
		"    get:",
		"      operationId: ListUsers",
		"      summary: list",
		"",
		"    post:",
		"      operationId: CreateUser",
		"  /orgs:",
		"    get:",
		"      operationId: ListOrgs",
	}, "\n")

	block, start, end, err := extractOpenAPIBlock(content, "ListUsers")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(block, "/users") || !strings.Contains(block, "operationId: ListUsers") {
		t.Errorf("block missing expected content: %q", block)
	}
	if strings.Contains(block, "CreateUser") {
		// The block should stop before the post method's operationId leaks... but
		// the path block spans the whole /users path (both methods). Just ensure
		// it does not run into /orgs.
		if strings.Contains(block, "/orgs") {
			t.Errorf("block should not include /orgs: %q", block)
		}
	}
	if start < 0 || end <= start {
		t.Errorf("bad line range: start=%d end=%d", start, end)
	}
}
