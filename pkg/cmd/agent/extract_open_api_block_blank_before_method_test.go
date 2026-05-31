//ff:func feature=agent type=test control=sequence
//ff:what TestExtractOpenAPIBlock — op 추출 성공 + op미존재/method미존재/path미존재 에러 분기 검증
package agent

import (
	"strings"
	"testing"
)

func TestExtractOpenAPIBlockBlankBeforeMethod(t *testing.T) {
	// A blank line between the path key and the HTTP method exercises the
	// empty-line continue inside the path-line search loop.
	content := strings.Join([]string{
		"paths:",
		"  /users:",
		"",
		"    get:",
		"      operationId: ListUsers",
	}, "\n")
	block, _, _, err := extractOpenAPIBlock(content, "ListUsers")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(block, "/users") {
		t.Errorf("block missing path: %q", block)
	}
}
