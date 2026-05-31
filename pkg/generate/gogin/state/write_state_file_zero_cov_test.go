//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what TestBuildTransitionMap_ZeroCov — 초기 전이 제외 + 중첩 맵 구성
package state

import (
	"os"
	"path/filepath"
	"testing"
)

func TestWriteStateFile_ZeroCov(t *testing.T) {
	dir := t.TempDir()
	transMap := map[string]map[string]string{"draft": {"submit": "review"}}
	if err := writeStateFile(dir, "course", "Course", transMap); err != nil {
		t.Fatalf("writeStateFile error: %v", err)
	}
	for _, name := range []string{"course.go", "course_can_transition.go", "course_next_state.go"} {
		if _, err := os.Stat(filepath.Join(dir, name)); err != nil {
			t.Errorf("expected file %s: %v", name, err)
		}
	}
}
