//ff:func feature=agent type=test control=sequence
//ff:what TestSearchDocs — 미지원 레이어는 빈 결과, 키워드 매칭 시 docs 섹션 반환 검증
package agent

import (
	"strings"
	"testing"
)

func TestSearchDocsRuleIDMatch(t *testing.T) {
	// ddl.md references the rule ID "XDO-77"; a diagnostic carrying that rule ID
	// in bracket form drives the rule-ID extraction (Phase 1) and the section
	// match by rule ID, taking priority over keyword matching.
	got := searchDocs(layerDDL, []string{"[XDO-77] uuid format mismatch"})
	if got == "" {
		t.Fatal("expected a docs section for XDO-77 rule-id match")
	}
	if !strings.Contains(got, "XDO-77") {
		t.Errorf("matched section missing rule id: %q", got)
	}

	// Two diagnostics carrying the same rule ID: the first marks the section as
	// seen, the second hits the seen[i] continue guard inside the Phase-1 loop.
	dup := searchDocs(layerDDL, []string{"[XDO-77] first", "[XDO-77] second"})
	if !strings.Contains(dup, "XDO-77") {
		t.Errorf("dup rule-id match missing section: %q", dup)
	}
}
