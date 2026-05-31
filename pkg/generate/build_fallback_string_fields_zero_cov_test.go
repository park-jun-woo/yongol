//ff:func feature=generate type=test control=sequence
//ff:what TestGenerateBatch_ZeroCov — pkg/generate 소형 순수/IO 헬퍼 분기 커버
package generate

import (
	"testing"
)

func TestBuildFallbackStringFields_ZeroCov(t *testing.T) {
	got := buildFallbackStringFields([]string{"a", "b"})
	if len(got) != 2 || got["a"].Type != "string" {
		t.Fatalf("buildFallbackStringFields = %v", got)
	}
}
