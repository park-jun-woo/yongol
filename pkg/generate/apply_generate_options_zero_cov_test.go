//ff:func feature=generate type=test control=sequence
//ff:what TestGenerateBatch_ZeroCov — pkg/generate 소형 순수/IO 헬퍼 분기 커버
package generate

import (
	"testing"
)

func TestApplyGenerateOptions_ZeroCov(t *testing.T) {
	cfg := &generateConfig{}
	called := false
	applyGenerateOptions(cfg, []GenerateOption{func(c *generateConfig) { called = true }})
	if !called {
		t.Error("expected hook called")
	}
	// empty hooks → no-op
	applyGenerateOptions(cfg, nil)
}
