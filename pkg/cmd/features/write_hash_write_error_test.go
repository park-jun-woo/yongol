//ff:func feature=features type=test control=sequence
//ff:what TestWriteHash — .yongol 해시 기록 성공 / 디렉토리 부재 시 write 에러 분기 검증
package features

import (
	"path/filepath"
	"strings"
	"testing"
)

func TestWriteHash_WriteError(t *testing.T) {
	// specsDir doesn't exist -> WriteFile fails.
	missing := filepath.Join(t.TempDir(), "no", "such", "dir")
	if err := writeHash(missing, []byte("x")); err == nil ||
		!strings.Contains(err.Error(), "write .yongol") {
		t.Fatalf("want write error, got %v", err)
	}
}
