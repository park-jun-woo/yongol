//ff:func feature=gen-react type=test control=iteration dimension=1
//ff:what 단순 react 헬퍼 (pick*/sorted*/orDefault/quotedUnion/naivePluralize) 묶음 커버 — 같은 basename 미사용 함수용
package react

import (
	"testing"
)

func TestNaivePluralizeZC(t *testing.T) {
	cases := map[string]string{
		"": "", "course": "courses", "class": "classes",
		"dish": "dishes", "match": "matches", "box": "boxes", "quiz": "quizes",
	}
	for in, want := range cases {
		if got := naivePluralize(in); got != want {
			t.Errorf("naivePluralize(%q) = %q, want %q", in, got, want)
		}
	}
}
