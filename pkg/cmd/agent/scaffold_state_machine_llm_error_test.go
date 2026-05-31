//ff:func feature=agent type=test control=sequence
//ff:what TestScaffoldStateMachine — states 없음 0 / 기존파일 skip / mkdir 에러 / 미지원 backend LLM 에러 분기 검증
package agent

import (
	"bytes"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

func TestScaffoldStateMachineLLMError(t *testing.T) {
	// Missing states file + unsupported backend -> target's LLM call fails, propagated.
	ff := &features.FeaturesFile{Tables: map[string]features.TableDef{
		"orders": {States: []string{"pending"}},
	}}
	var out bytes.Buffer
	cfg := Config{Backend: "unsupported-backend", Model: "m"}
	if _, err := scaffoldStateMachine(t.TempDir(), ff, nil, cfg, &out); err == nil {
		t.Fatal("expected LLM error from scaffoldStateMachineTarget")
	}
}
