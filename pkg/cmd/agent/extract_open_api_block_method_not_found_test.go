//ff:func feature=agent type=test control=sequence
//ff:what TestExtractOpenAPIBlock — op 추출 성공 + op미존재/method미존재/path미존재 에러 분기 검증
package agent

import (
	"strings"
	"testing"
)

func TestExtractOpenAPIBlockMethodNotFound(t *testing.T) {
	// operationId with no HTTP method line above it.
	content := "info:\n  operationId: Weird\n"
	_, _, _, err := extractOpenAPIBlock(content, "Weird")
	if err == nil || !strings.Contains(err.Error(), "method line") {
		t.Fatalf("expected method-line error, got: %v", err)
	}
}
