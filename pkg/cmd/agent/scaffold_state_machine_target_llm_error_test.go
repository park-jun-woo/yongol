//ff:func feature=agent type=test control=sequence
//ff:what TestScaffoldStateMachine — states 없음 0 / 기존파일 skip / mkdir 에러 / 미지원 backend LLM 에러 분기 검증
package agent

import (
	"bytes"
	"testing"
)

func TestScaffoldStateMachineTargetLLMError(t *testing.T) {
	var out bytes.Buffer
	cfg := Config{Backend: "unsupported-backend", Model: "m"}
	created, err := scaffoldStateMachineTarget(t.TempDir(), "orders", []string{"pending"}, nil, "sys", cfg, &out)
	if err == nil {
		t.Fatal("expected LLM error")
	}
	if created {
		t.Fatal("expected created=false on error")
	}
}
