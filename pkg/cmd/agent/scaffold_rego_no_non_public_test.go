//ff:func feature=agent type=test control=sequence
//ff:what TestScaffoldRego — 기존파일 skip / non-public 없음 skip / non-public 존재+미지원 backend LLM 에러 분기 검증
package agent

import (
	"bytes"
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

func TestScaffoldRegoNoNonPublic(t *testing.T) {
	// Only public features → no non-public → skip.
	ff := &features.FeaturesFile{Features: []features.Feature{{Op: "Ping", Public: true}}}
	var out bytes.Buffer
	if err := scaffoldRego(t.TempDir(), ff, nil, Config{}, &out); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(out.String(), "no non-public") {
		t.Errorf("expected no-non-public message, got: %q", out.String())
	}
}
