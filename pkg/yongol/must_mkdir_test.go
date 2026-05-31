//ff:func feature=orchestrator type=test control=sequence
//ff:what findYongolPkgRoot — env 미설정 시 CWD fallback이 sibling ssac/pkg를 반환
package yongol

import (
	"os"
	"testing"
)

func mustMkdir(t *testing.T, dir string) {
	t.Helper()
	if err := os.MkdirAll(dir, 0o755); err != nil {
		t.Fatal(err)
	}
}
