//ff:func feature=orchestrator type=test control=sequence
//ff:what assertSelectColumns — extractSelectColumns 결과를 기대값과 비교하는 테스트 헬퍼

package sqlc

import (
	"strings"
	"testing"
)

// assertSelectColumns compares extractSelectColumns output against expected values.
func assertSelectColumns(t *testing.T, star bool, cols []string, wantStar bool, wantCols []string) {
	t.Helper()
	if star != wantStar {
		t.Errorf("selectStar = %v, want %v", star, wantStar)
	}
	if wantCols == nil {
		if cols != nil {
			t.Errorf("selectCols = %v, want nil", cols)
		}
	} else {
		gotStr := strings.Join(cols, ",")
		wantStr := strings.Join(wantCols, ",")
		if gotStr != wantStr {
			t.Errorf("selectCols = %v, want %v", cols, wantCols)
		}
	}
}
