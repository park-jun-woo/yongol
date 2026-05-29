//ff:func feature=agent type=test control=selection dimension=1
//ff:what TestBuildErrorResponses — public/delete/기본 분기별 에러 응답 코드 집합 검증

package agent

import (
	"sort"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

func TestBuildErrorResponses(t *testing.T) {
	codes := func(m map[string]any) []string {
		out := make([]string, 0, len(m))
		for k := range m {
			out = append(out, k)
		}
		sort.Strings(out)
		return out
	}
	eq := func(a, b []string) bool {
		if len(a) != len(b) {
			return false
		}
		for i := range a {
			if a[i] != b[i] {
				return false
			}
		}
		return true
	}

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
			got := buildErrorResponses(tc.feat)
			if gc := codes(got); !eq(gc, tc.want) {
				t.Errorf("codes: got %v, want %v", gc, tc.want)
			}
			// Each response must carry the Error $ref schema (copyMap produced a value).
			for code, resp := range got {
				rm, ok := resp.(map[string]any)
				if !ok {
					t.Fatalf("response %s is not a map: %v", code, resp)
				}
				if rm["description"] != "Error" {
					t.Errorf("response %s missing Error description", code)
				}
			}
		})
	}
}
