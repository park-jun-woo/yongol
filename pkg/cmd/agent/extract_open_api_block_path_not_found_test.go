//ff:func feature=agent type=test control=sequence
//ff:what TestExtractOpenAPIBlock — op 추출 성공 + op미존재/method미존재/path미존재 에러 분기 검증
package agent

import (
	"strings"
	"testing"
)

func TestExtractOpenAPIBlockPathNotFound(t *testing.T) {
	// HTTP method at column 0 (indent 0): no line with smaller indent exists, so
	// the path line cannot be found.
	content := "get:\n  operationId: Flat\n"
	_, _, _, err := extractOpenAPIBlock(content, "Flat")
	if err == nil || !strings.Contains(err.Error(), "path line") {
		t.Fatalf("expected path-line error, got: %v", err)
	}
}
