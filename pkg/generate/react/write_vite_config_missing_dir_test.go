//ff:func feature=gen-react type=test control=sequence
//ff:what writeViteConfig vite.config.ts 생성 내용·에러경로 검증
package react

import (
	"path/filepath"
	"testing"
)

func TestWriteViteConfigMissingDir(t *testing.T) {
	missing := filepath.Join(t.TempDir(), "absent", "dir")
	if err := writeViteConfig(missing); err == nil {
		t.Fatal("expected error writing into non-existent directory, got nil")
	}
}
