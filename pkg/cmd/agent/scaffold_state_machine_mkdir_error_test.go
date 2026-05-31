//ff:func feature=agent type=test control=sequence
//ff:what TestScaffoldStateMachine — states 없음 0 / 기존파일 skip / mkdir 에러 / 미지원 backend LLM 에러 분기 검증
package agent

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

func TestScaffoldStateMachineMkdirError(t *testing.T) {
	// Creating "states" as a regular file makes os.MkdirAll(statesDir) fail.
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "states"), []byte("x"), 0o644); err != nil {
		t.Fatal(err)
	}
	ff := &features.FeaturesFile{Tables: map[string]features.TableDef{
		"orders": {States: []string{"pending"}},
	}}
	var out bytes.Buffer
	if _, err := scaffoldStateMachine(dir, ff, nil, Config{}, &out); err == nil {
		t.Fatal("expected mkdir error when states is a file")
	}
}
