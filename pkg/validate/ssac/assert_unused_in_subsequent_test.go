//ff:func feature=validate type=test-helper control=sequence topic=ssac-structural
//ff:what assertUnusedInSubsequent — s62unusedInSubsequent 결과 일치 검증 헬퍼

package ssac

import (
	"testing"

	parsessac "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func assertUnusedInSubsequent(t *testing.T, varName string, seqs []parsessac.Sequence, start int, want bool) {
	t.Helper()
	got := s62unusedInSubsequent(varName, seqs, start)
	if got != want {
		t.Errorf("s62unusedInSubsequent(%q, seqs, %d) = %v, want %v",
			varName, start, got, want)
	}
}
