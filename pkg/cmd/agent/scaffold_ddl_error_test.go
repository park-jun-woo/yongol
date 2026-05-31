//ff:func feature=agent type=test control=sequence
//ff:what TestScaffold — 테이블 없음 skip / 테이블 존재+미지원 backend → DDL phase 에러 분기 검증
package agent

import (
	"bytes"
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

func TestScaffoldDDLError(t *testing.T) {
	// One table + an unsupported backend makes the DDL phase's LLM call fail, so
	// scaffold returns a wrapped "scaffold DDL" error.
	ff := &features.FeaturesFile{
		Tables: map[string]features.TableDef{
			"users": {},
		},
	}
	var out bytes.Buffer
	cfg := Config{Backend: "unsupported-backend", Model: "m"}
	err := scaffold(t.TempDir(), ff, nil, cfg, &out)
	if err == nil {
		t.Fatal("expected DDL phase error")
	}
	if !strings.Contains(err.Error(), "scaffold DDL") {
		t.Errorf("expected scaffold DDL error, got: %v", err)
	}
}
