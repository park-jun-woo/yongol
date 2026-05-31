//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestSortedColumnKeys — 컬럼 맵 키를 오름차순 정렬 반환 + 빈 맵 처리 검증
package sqlcpost

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

func TestSortedColumnKeys_Empty(t *testing.T) {
	if got := sortedColumnKeys(map[string]ddl.Column{}); len(got) != 0 {
		t.Errorf("expected empty result, got %v", got)
	}
}
