//ff:func feature=features type=test control=sequence
//ff:what TestRunAdd — new/existing 파싱 에러 / no-new-diff / 신규 op stub 생성 + skip 기존파일 분기 검증
package features

import (
	"bytes"
	"path/filepath"
	"strings"
	"testing"
)

func TestRunAdd_ExistingParseError(t *testing.T) {
	dir := t.TempDir()
	newPath := filepath.Join(dir, "new.yaml")
	writeFeats(t, newPath, baseFeats)
	// specsDir has no features.yaml -> existing load fails.
	var out bytes.Buffer
	if err := RunAdd(&out, dir, newPath); err == nil ||
		!strings.Contains(err.Error(), "existing features") {
		t.Fatalf("want existing features error, got %v", err)
	}
}
