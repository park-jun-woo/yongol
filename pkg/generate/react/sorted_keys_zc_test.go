//ff:func feature=gen-react type=test control=sequence
//ff:what 단순 react 헬퍼 (pick*/sorted*/orDefault/quotedUnion/naivePluralize) 묶음 커버 — 같은 basename 미사용 함수용
package react

import (
	"reflect"
	"testing"
)

func TestSortedKeysZC(t *testing.T) {
	if got := sortedKeys(map[string]string{"b": "1", "a": "2"}); !reflect.DeepEqual(got, []string{"a", "b"}) {
		t.Errorf("sortedKeys = %v", got)
	}
}
