//ff:func feature=agent type=test control=sequence
//ff:what TestScaffoldDDL — 테이블 없음 nil / 기존파일 skip 성공 / 미지원 backend LLM 에러 분기 검증
package agent

import (
	"bytes"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

func TestScaffoldDDLNoTables(t *testing.T) {
	var out bytes.Buffer
	if err := scaffoldDDL(t.TempDir(), &features.FeaturesFile{}, nil, Config{}, &out); err != nil {
		t.Fatalf("no tables: unexpected error: %v", err)
	}
}
