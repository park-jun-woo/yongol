//ff:func feature=gen-react type=test control=sequence
//ff:what writeTSConfig tsconfig.json 생성 내용·유효 JSON·에러경로 검증
package react

import (
	"path/filepath"
	"testing"
)

func TestWriteTSConfigMissingDir(t *testing.T) {
	missing := filepath.Join(t.TempDir(), "absent", "dir")
	if err := writeTSConfig(missing); err == nil {
		t.Fatal("expected error writing into non-existent directory, got nil")
	}
}
