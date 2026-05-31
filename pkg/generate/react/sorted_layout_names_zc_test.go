//ff:func feature=gen-react type=test control=sequence
//ff:what 단순 react 헬퍼 (pick*/sorted*/orDefault/quotedUnion/naivePluralize) 묶음 커버 — 같은 basename 미사용 함수용
package react

import (
	"testing"
)

func TestSortedLayoutNamesZC(t *testing.T) {
	grouped := map[string][]stmlRoute{"": nil, "app": nil, "auth": nil}
	got := sortedLayoutNames(grouped)
	if len(got) != 3 || got[0] != "app" || got[1] != "auth" || got[2] != "" {
		t.Errorf("sortedLayoutNames = %v", got)
	}
}
