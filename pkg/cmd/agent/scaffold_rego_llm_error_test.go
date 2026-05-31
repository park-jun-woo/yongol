//ff:func feature=agent type=test control=sequence
//ff:what TestScaffoldRego — 기존파일 skip / non-public 없음 skip / non-public 존재+미지원 backend LLM 에러 분기 검증
package agent

import (
	"bytes"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

func TestScaffoldRegoLLMError(t *testing.T) {
	// A non-public feature triggers the batch LLM call, which fails for the
	// unsupported backend.
	ff := &features.FeaturesFile{Features: []features.Feature{{Op: "CreateUser", Public: false}}}
	var out bytes.Buffer
	cfg := Config{Backend: "unsupported-backend", Model: "m"}
	if err := scaffoldRego(t.TempDir(), ff, nil, cfg, &out); err == nil {
		t.Fatal("expected LLM error from rego batch")
	}
}
