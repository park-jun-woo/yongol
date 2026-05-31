//ff:func feature=agent type=test control=sequence
//ff:what TestScaffoldDDL — 테이블 없음 nil / 기존파일 skip 성공 / 미지원 backend LLM 에러 분기 검증
package agent

import (
	"bytes"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

func TestScaffoldDDLLLMError(t *testing.T) {
	// A missing DDL file + unsupported backend makes scaffoldDDLTable's LLM call
	// fail, propagated out of scaffoldDDL.
	ff := &features.FeaturesFile{Tables: map[string]features.TableDef{"users": {}}}
	var out bytes.Buffer
	cfg := Config{Backend: "unsupported-backend", Model: "m"}
	if err := scaffoldDDL(t.TempDir(), ff, nil, cfg, &out); err == nil {
		t.Fatal("expected LLM error from scaffoldDDLTable")
	}
}
