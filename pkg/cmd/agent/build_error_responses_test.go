//ff:func feature=agent type=test control=iteration dimension=1
//ff:what TestBuildErrorResponses — public/delete/기본 분기별 에러 응답 코드 집합 검증
package agent

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

func TestBuildErrorResponses(t *testing.T) {
	cases := []struct {
		name string
		feat features.Feature
		want []string
	}{
		{"public", features.Feature{Op: "ListThings", Public: true}, []string{"400", "500"}},
		{"private delete", features.Feature{Op: "DeleteThing", Public: false}, []string{"401", "403", "404"}},
		{"private default", features.Feature{Op: "CreateThing", Public: false}, []string{"401", "403", "404", "500"}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			assertBuildErrorResponses(t, tc.feat, tc.want)
		})
	}
}
