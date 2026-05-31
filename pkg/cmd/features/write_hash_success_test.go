//ff:func feature=features type=test control=sequence
//ff:what TestWriteHash — .yongol 해시 기록 성공 / 디렉토리 부재 시 write 에러 분기 검증
package features

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestWriteHash_Success(t *testing.T) {
	specs := t.TempDir()
	if err := writeHash(specs, []byte("features: []\n")); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	got, err := os.ReadFile(filepath.Join(specs, ".yongol"))
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(got), "features.yaml: sha256:") {
		t.Errorf("unexpected .yongol content: %q", got)
	}
}
