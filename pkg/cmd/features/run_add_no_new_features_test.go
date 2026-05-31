//ff:func feature=features type=test control=sequence
//ff:what TestRunAdd — new/existing 파싱 에러 / no-new-diff / 신규 op stub 생성 + skip 기존파일 분기 검증
package features

import (
	"bytes"
	"path/filepath"
	"strings"
	"testing"
)

func TestRunAdd_NoNewFeatures(t *testing.T) {
	specs := t.TempDir()
	writeFeats(t, filepath.Join(specs, "features.yaml"), baseFeats)
	newPath := filepath.Join(t.TempDir(), "new.yaml")
	writeFeats(t, newPath, baseFeats) // identical ops -> no diff
	var out bytes.Buffer
	if err := RunAdd(&out, specs, newPath); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(out.String(), "no new features") {
		t.Errorf("want no-new-features message, got %q", out.String())
	}
}
