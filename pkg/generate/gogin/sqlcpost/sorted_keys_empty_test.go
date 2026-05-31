//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestSortedKeys — string 맵 키를 오름차순 정렬 반환 + 빈 맵 처리 검증
package sqlcpost

import (
	"testing"
)

func TestSortedKeys_Empty(t *testing.T) {
	if got := sortedKeys(map[string]string{}); len(got) != 0 {
		t.Errorf("expected empty result, got %v", got)
	}
}
