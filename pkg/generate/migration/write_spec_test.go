//ff:func feature=migration type=test-helper control=sequence
//ff:what writeSpec — 테스트용 DDL 파일 생성 유틸
package migration

import (
	"os"
	"path/filepath"
	"testing"
)

func writeSpec(t *testing.T, dir, name, body string) {
	t.Helper()
	if err := os.MkdirAll(dir, 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, name), []byte(body), 0644); err != nil {
		t.Fatal(err)
	}
}
