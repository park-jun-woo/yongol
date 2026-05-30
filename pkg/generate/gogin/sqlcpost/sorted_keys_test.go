//ff:func feature=gen-gogin type=test control=iteration topic=sqlc-post
//ff:what TestSortedKeys — string 맵 키를 오름차순 정렬 반환 + 빈 맵 처리 검증

package sqlcpost

import (
	"reflect"
	"testing"
)

func TestSortedKeys(t *testing.T) {
	got := sortedKeys(map[string]string{"c": "3", "a": "1", "b": "2"})
	want := []string{"a", "b", "c"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("sortedKeys = %v, want %v", got, want)
	}
}

func TestSortedKeys_Empty(t *testing.T) {
	if got := sortedKeys(map[string]string{}); len(got) != 0 {
		t.Errorf("expected empty result, got %v", got)
	}
}
