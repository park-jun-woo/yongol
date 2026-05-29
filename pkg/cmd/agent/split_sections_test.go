//ff:func feature=agent type=test control=iteration dimension=2
//ff:what TestSplitSections — markdown 콘텐츠를 ## 헤딩별 섹션으로 분리하는지 검증

package agent

import (
	"strings"
	"testing"
)

func TestSplitSections(t *testing.T) {
	content := "intro line ignored\n## First\nbody1\n## Second\nbody2\nmore"
	sections := splitSections(content)
	if len(sections) != 2 {
		t.Fatalf("sections = %d, want 2", len(sections))
	}
	if !strings.HasPrefix(sections[0], "## First") || !strings.Contains(sections[0], "body1") {
		t.Errorf("section0 = %q", sections[0])
	}
	if !strings.HasPrefix(sections[1], "## Second") || !strings.Contains(sections[1], "more") {
		t.Errorf("section1 = %q", sections[1])
	}
	if strings.Contains(sections[0], "## Second") {
		t.Errorf("section0 leaked into section1: %q", sections[0])
	}

	if got := splitSections("no headings here"); got != nil {
		t.Errorf("content without headings = %v, want nil", got)
	}
}
