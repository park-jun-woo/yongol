//ff:func feature=features type=test control=sequence
//ff:what TestRunRemove — 빈 ops/존재안함/abort/확인삭제/ssac-skip/--yes 성공 분기 검증
package features

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestRunRemove_Abort(t *testing.T) {
	specs := setupSpecs(t)
	var out bytes.Buffer
	// Answer "n" -> aborted, features.yaml untouched.
	if err := RunRemove(&out, strings.NewReader("n\n"), specs, []string{"DeleteTask"}, false); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(out.String(), "aborted") {
		t.Errorf("want aborted message, got %q", out.String())
	}
	got, _ := os.ReadFile(filepath.Join(specs, "features.yaml"))
	if !strings.Contains(string(got), "DeleteTask") {
		t.Error("features.yaml should be untouched after abort")
	}
}
