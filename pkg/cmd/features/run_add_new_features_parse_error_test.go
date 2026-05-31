//ff:func feature=features type=test control=sequence
//ff:what TestRunAdd — new/existing 파싱 에러 / no-new-diff / 신규 op stub 생성 + skip 기존파일 분기 검증
package features

import (
	"bytes"
	"path/filepath"
	"strings"
	"testing"
)

func TestRunAdd_NewFeaturesParseError(t *testing.T) {
	var out bytes.Buffer
	if err := RunAdd(&out, t.TempDir(), filepath.Join(t.TempDir(), "missing.yaml")); err == nil ||
		!strings.Contains(err.Error(), "new features") {
		t.Fatalf("want new features error, got %v", err)
	}
}
