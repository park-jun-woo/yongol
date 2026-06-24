//ff:func feature=gen-react type=test control=sequence
//ff:what 테스트 헬퍼 — 경로의 파일 내용을 문자열로 읽고 실패 시 Fatal

package react

import (
	"os"
	"testing"
)

func readFile(t *testing.T, path string) string {
	t.Helper()
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read %s: %v", path, err)
	}
	return string(data)
}
