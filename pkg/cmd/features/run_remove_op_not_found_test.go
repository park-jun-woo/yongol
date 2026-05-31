//ff:func feature=features type=test control=sequence
//ff:what TestRunRemove — 빈 ops/존재안함/abort/확인삭제/ssac-skip/--yes 성공 분기 검증
package features

import (
	"bytes"
	"strings"
	"testing"
)

func TestRunRemove_OpNotFound(t *testing.T) {
	specs := setupSpecs(t)
	var out bytes.Buffer
	if err := RunRemove(&out, strings.NewReader(""), specs, []string{"Ghost"}, true); err == nil ||
		!strings.Contains(err.Error(), "not found") {
		t.Fatalf("want not-found error, got %v", err)
	}
}
