//ff:func feature=agent type=test control=sequence
//ff:what TestScaffoldStateMachine — states 없음 0 / 기존파일 skip / mkdir 에러 / 미지원 backend LLM 에러 분기 검증
package agent

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

func TestScaffoldStateMachineSkipExisting(t *testing.T) {
	dir := t.TempDir()
	statesDir := filepath.Join(dir, "states")
	if err := os.MkdirAll(statesDir, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(statesDir, "orders.md"), []byte("stateDiagram\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	ff := &features.FeaturesFile{Tables: map[string]features.TableDef{
		"orders": {States: []string{"pending", "shipped"}},
	}}
	var out bytes.Buffer
	n, err := scaffoldStateMachine(dir, ff, nil, Config{}, &out)
	if err != nil {
		t.Fatalf("skip-existing: unexpected error: %v", err)
	}
	if n != 0 {
		t.Fatalf("skip-existing: expected count 0, got %d", n)
	}
	if !strings.Contains(out.String(), "skipped") {
		t.Errorf("expected skip message, got: %q", out.String())
	}
}
