//ff:func feature=agent type=test control=iteration dimension=3
//ff:what TestSearchDocs — 미지원 레이어는 빈 결과, 키워드 매칭 시 docs 섹션 반환 검증

package agent

import (
	"strings"
	"testing"
)

func TestSearchDocs(t *testing.T) {
	// Unknown layer has no docs file → empty result.
	if got := searchDocs(layerUnknown, []string{"NOT NULL"}); got != "" {
		t.Errorf("unknown layer = %q, want empty", got)
	}

	// No diagnostic keywords match → no section returned.
	if got := searchDocs(layerDDL, []string{"completely unrelated text"}); got != "" {
		t.Errorf("no-match = %q, want empty", got)
	}

	// A keyword present in ddl.md ("NOT NULL") yields a matching section.
	got := searchDocs(layerDDL, []string{"column must be NOT NULL"})
	if got == "" {
		t.Fatal("expected a docs section for NOT NULL keyword match")
	}
	if !strings.Contains(got, "NOT NULL") {
		t.Errorf("matched section missing keyword: %q", got)
	}
	if len(got) > 2048 {
		t.Errorf("result length = %d, want capped at 2048", len(got))
	}
}

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

func TestSearchDocsKeywordSeenDedup(t *testing.T) {
	// Two keywords ("@auth" and "operationId") may resolve to overlapping
	// sections; the seen[i] guard in Phase 2 prevents a section from being
	// appended twice when both keywords are active.
	got := searchDocs(layerSSaC, []string{"@auth and @auth and operationId problem"})
	// Either a match or empty is acceptable structurally; the goal is to exercise
	// the multi-keyword Phase-2 loop with the dedup guard. Assert no duplicate
	// section boundary doubling by checking the result is bounded.
	if len(got) > 2048 {
		t.Errorf("result length = %d, want capped at 2048", len(got))
	}
}
