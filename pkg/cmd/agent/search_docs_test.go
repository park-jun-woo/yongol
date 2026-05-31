//ff:func feature=agent type=test control=sequence
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
