//ff:func feature=features type=test control=sequence
//ff:what TestRunAdd — new/existing 파싱 에러 / no-new-diff / 신규 op stub 생성 + skip 기존파일 분기 검증
package features

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestRunAdd_SkipExistingStub(t *testing.T) {
	specs := t.TempDir()
	writeFeats(t, filepath.Join(specs, "features.yaml"), baseFeats)
	// Pre-create the stub so the new op's write is skipped.
	stubDir := filepath.Join(specs, "service", "task")
	if err := os.MkdirAll(stubDir, 0o755); err != nil {
		t.Fatal(err)
	}
	writeFeats(t, filepath.Join(stubDir, "DeleteTask.ssac"), "existing")
	newContent := baseFeats + `  - op: DeleteTask
    path: DELETE /tasks/{id}
    desc: Delete a task
`
	newPath := filepath.Join(t.TempDir(), "new.yaml")
	writeFeats(t, newPath, newContent)

	var out bytes.Buffer
	if err := RunAdd(&out, specs, newPath); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(out.String(), "skip") {
		t.Errorf("want skip message, got %q", out.String())
	}
	// Pre-existing stub content preserved.
	got, _ := os.ReadFile(filepath.Join(stubDir, "DeleteTask.ssac"))
	if string(got) != "existing" {
		t.Errorf("stub overwritten: %q", got)
	}
}
