//ff:func feature=gen-react type=test control=sequence
//ff:what 테스트 헬퍼 — 부모 디렉토리 생성 후 OpenAPI spec 파일 기록

package react

import (
	"os"
	"path/filepath"
	"testing"
)

func writeSpecFile(t *testing.T, path, body string) {
	t.Helper()
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatalf("mkdir %s: %v", filepath.Dir(path), err)
	}
	if err := os.WriteFile(path, []byte(body), 0o644); err != nil {
		t.Fatalf("write %s: %v", path, err)
	}
}
