//ff:func feature=agent type=test control=sequence
//ff:what TestChainResolve — SSOT 미검출/파싱 진단 시 빈 반환 + 예제 specs 로 link 매칭 성공 경로 검증
package agent

import (
	"os"
	"path/filepath"
	"testing"
)

func TestChainResolveDetectError(t *testing.T) {
	// Pointing specsDir at a regular file (not a directory) makes DetectSSOTs
	// fail, exercising the early "" "" return at the top of chainResolve.
	file := filepath.Join(t.TempDir(), "not-a-dir.txt")
	if err := os.WriteFile(file, []byte("x"), 0o644); err != nil {
		t.Fatal(err)
	}
	desc, path := chainResolve(file, "whatever", nil)
	if desc != "" || path != "" {
		t.Errorf("chainResolve on file = %q, %q; want empty, empty", desc, path)
	}
}
