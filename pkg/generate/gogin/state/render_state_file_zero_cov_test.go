//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what TestBuildTransitionMap_ZeroCov — 초기 전이 제외 + 중첩 맵 구성
package state

import (
	"strings"
	"testing"
)

func TestRenderStateFile_ZeroCov(t *testing.T) {
	src := renderStateFile("course", "Course", map[string]map[string]string{
		"draft": {"submit": "review"},
	})
	for _, want := range []string{"package statemachine", "var CourseTransitions", "states/course.md"} {
		if !strings.Contains(src, want) {
			t.Errorf("renderStateFile missing %q", want)
		}
	}
}
