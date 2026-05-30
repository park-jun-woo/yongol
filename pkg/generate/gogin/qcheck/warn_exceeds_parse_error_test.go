//ff:func feature=gen-gogin type=test control=branch topic=depth-report
//ff:what TestWarnExceeds_ParseError — 파서 에러 시 WARN 단일 메시지 반환 검증

package qcheck

import (
	"strings"
	"testing"
)

func TestWarnExceeds_ParseError(t *testing.T) {
	warns := WarnExceeds("bad.go", "@@@ not go", DefaultLimits())
	if len(warns) != 1 || !strings.Contains(warns[0], "parse error") {
		t.Fatalf("expected single parse-error WARN, got %v", warns)
	}
}
