//ff:func feature=features type=test control=sequence
//ff:what TestRunRemove — 빈 ops/존재안함/abort/확인삭제/ssac-skip/--yes 성공 분기 검증
package features

import (
	"bytes"
	"strings"
	"testing"
)

func TestRunRemove_ExistingLoadError(t *testing.T) {
	var out bytes.Buffer
	if err := RunRemove(&out, strings.NewReader(""), t.TempDir(), []string{"X"}, true); err == nil ||
		!strings.Contains(err.Error(), "existing features") {
		t.Fatalf("want existing features error, got %v", err)
	}
}
