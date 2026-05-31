//ff:func feature=agent type=test-helper control=iteration dimension=1
//ff:what assertBuildErrorResponses — buildErrorResponses 결과 코드 집합 + Error description 검증 헬퍼
package agent

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

// assertBuildErrorResponses asserts that buildErrorResponses(feat) yields the
// expected status codes, each carrying the Error $ref schema.
func assertBuildErrorResponses(t *testing.T, feat features.Feature, want []string) {
	t.Helper()
	got := buildErrorResponses(feat)
	if gc := errorResponseCodes(got); !equalStrings(gc, want) {
		t.Errorf("codes: got %v, want %v", gc, want)
	}
	for code, resp := range got {
		rm, ok := resp.(map[string]any)
		if !ok {
			t.Fatalf("response %s is not a map: %v", code, resp)
		}
		if rm["description"] != "Error" {
			t.Errorf("response %s missing Error description", code)
		}
	}
}
