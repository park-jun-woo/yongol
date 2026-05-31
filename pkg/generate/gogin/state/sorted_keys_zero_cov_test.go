//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestBuildTransitionMap_ZeroCov — 초기 전이 제외 + 중첩 맵 구성
package state

import (
	"strings"
	"testing"
)

func TestSortedKeys_ZeroCov(t *testing.T) {
	m := map[string]int{"c": 1, "a": 2, "b": 3}
	got := sortedKeys(m)
	if strings.Join(got, ",") != "a,b,c" {
		t.Errorf("sortedKeys = %v", got)
	}
}
