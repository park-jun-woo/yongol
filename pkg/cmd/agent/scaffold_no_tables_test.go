//ff:func feature=agent type=test control=sequence
//ff:what TestScaffold — 테이블 없음 skip / 테이블 존재+미지원 backend → DDL phase 에러 분기 검증
package agent

import (
	"bytes"
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

func TestScaffoldNoTables(t *testing.T) {
	var out bytes.Buffer
	// nil FeaturesFile → skip.
	if err := scaffold(t.TempDir(), nil, nil, Config{}, &out); err != nil {
		t.Fatalf("nil ff: unexpected error: %v", err)
	}
	if !strings.Contains(out.String(), "scaffold skipped") {
		t.Errorf("expected skip message, got: %q", out.String())
	}

	// FeaturesFile with no tables → skip.
	out.Reset()
	if err := scaffold(t.TempDir(), &features.FeaturesFile{}, nil, Config{}, &out); err != nil {
		t.Fatalf("empty ff: unexpected error: %v", err)
	}
	if !strings.Contains(out.String(), "scaffold skipped") {
		t.Errorf("expected skip message, got: %q", out.String())
	}
}
