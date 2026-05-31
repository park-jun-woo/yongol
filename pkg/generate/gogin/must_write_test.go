//ff:func feature=gen-gogin type=test control=sequence
//ff:what mustWrite — 테스트용 파일 기록 헬퍼
package gogin

import (
	"os"
	"testing"
)

func mustWrite(t *testing.T, path, content string) {
	t.Helper()
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatalf("write %s: %v", path, err)
	}
}
