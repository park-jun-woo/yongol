//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestBuildTransitionMap_ZeroCov — 초기 전이 제외 + 중첩 맵 구성
package state

import (
	"strings"
	"testing"
)

func TestRenderTransitionEntries_ZeroCov(t *testing.T) {
	var b strings.Builder
	transMap := map[string]map[string]string{
		"draft":  {"submit": "review"},                        // single event branch
		"review": {"approve": "published", "reject": "draft"}, // multi event branch
	}
	renderTransitionEntries(&b, transMap)
	out := b.String()
	if !strings.Contains(out, `"draft": {"submit": "review"}`) {
		t.Errorf("single-event line missing:\n%s", out)
	}
	if !strings.Contains(out, `"review": {`) || !strings.Contains(out, `"approve": "published"`) {
		t.Errorf("multi-event block missing:\n%s", out)
	}
}
