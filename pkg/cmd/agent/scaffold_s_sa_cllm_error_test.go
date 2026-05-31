//ff:func feature=agent type=test control=sequence
//ff:what TestScaffoldSSaC — features 없음 0,nil / 기존파일 skip(count 0) / 미지원 backend LLM 에러 분기 검증
package agent

import (
	"bytes"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

func TestScaffoldSSaCLLMError(t *testing.T) {
	// A missing SSaC file + unsupported backend makes scaffoldSSaCFeature's LLM
	// call fail, propagated out of scaffoldSSaC.
	ff := &features.FeaturesFile{Features: []features.Feature{{Op: "Login", Path: "/auth/login"}}}
	var out bytes.Buffer
	cfg := Config{Backend: "unsupported-backend", Model: "m"}
	if _, err := scaffoldSSaC(t.TempDir(), ff, "", nil, cfg, &out); err == nil {
		t.Fatal("expected LLM error from scaffoldSSaCFeature")
	}
}
