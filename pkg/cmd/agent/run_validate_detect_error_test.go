//ff:func feature=agent type=test control=sequence
//ff:what TestRunValidate — DetectSSOTs 에러 / 파싱 진단 / 정상 validate 진단 수집 분기 검증
package agent

import (
	"os"
	"path/filepath"
	"testing"
)

func TestRunValidateDetectError(t *testing.T) {
	// A file path (not a directory) makes DetectSSOTs fail.
	file := filepath.Join(t.TempDir(), "f.txt")
	if err := os.WriteFile(file, []byte("x"), 0o644); err != nil {
		t.Fatal(err)
	}
	if _, err := runValidate(file); err == nil {
		t.Fatal("expected detect error for non-directory path")
	}
}
