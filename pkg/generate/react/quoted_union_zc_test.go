//ff:func feature=gen-react type=test control=sequence
//ff:what 단순 react 헬퍼 (pick*/sorted*/orDefault/quotedUnion/naivePluralize) 묶음 커버 — 같은 basename 미사용 함수용
package react

import (
	"testing"
)

func TestQuotedUnionZC(t *testing.T) {
	if got := quotedUnion([]string{"a", "b"}); got != "'a' | 'b'" {
		t.Errorf("quotedUnion = %q", got)
	}
	if got := quotedUnion(nil); got != "" {
		t.Errorf("quotedUnion(nil) = %q", got)
	}
}
