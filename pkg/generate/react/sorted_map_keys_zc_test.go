//ff:func feature=gen-react type=test control=sequence
//ff:what 단순 react 헬퍼 (pick*/sorted*/orDefault/quotedUnion/naivePluralize) 묶음 커버 — 같은 basename 미사용 함수용
package react

import (
	"reflect"
	"testing"
)

func TestSortedMapKeysZC(t *testing.T) {
	if got := sortedMapKeys(map[string]string{"y": "1", "x": "2"}); !reflect.DeepEqual(got, []string{"x", "y"}) {
		t.Errorf("sortedMapKeys = %v", got)
	}
}
