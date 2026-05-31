//ff:func feature=agent type=test control=sequence
//ff:what TestExtractOpenAPIBlock — op 추출 성공 + op미존재/method미존재/path미존재 에러 분기 검증
package agent

import (
	"strings"
	"testing"
)

func TestExtractOpenAPIBlockOpNotFound(t *testing.T) {
	_, _, _, err := extractOpenAPIBlock("paths:\n  /x:\n    get:\n      operationId: Other\n", "Missing")
	if err == nil || !strings.Contains(err.Error(), "not found") {
		t.Fatalf("expected not-found error, got: %v", err)
	}
}
