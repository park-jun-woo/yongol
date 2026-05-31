//ff:func feature=report type=test control=sequence topic=sarif
//ff:what TestBuildDriverRules — catalog 전체 / fired-only fallback / 빈 fired nil 분기 검증
package sarif

import (
	"testing"
)

func TestBuildDriverRules_EmptyFired(t *testing.T) {
	if got := buildDriverRules(nil, map[string]struct{}{}); got != nil {
		t.Errorf("empty fired: got %+v, want nil", got)
	}
}
