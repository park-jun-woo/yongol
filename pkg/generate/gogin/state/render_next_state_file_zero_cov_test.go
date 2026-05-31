//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what TestBuildTransitionMap_ZeroCov — 초기 전이 제외 + 중첩 맵 구성
package state

import (
	"strings"
	"testing"
)

func TestRenderNextStateFile_ZeroCov(t *testing.T) {
	src := renderNextStateFile("course", "Course")
	for _, want := range []string{"package statemachine", "func CourseNextState", "CourseTransitions"} {
		if !strings.Contains(src, want) {
			t.Errorf("renderNextStateFile missing %q", want)
		}
	}
}
