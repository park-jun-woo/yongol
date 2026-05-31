//ff:func feature=agent type=test control=sequence
//ff:what TestScaffoldDDL — 테이블 없음 nil / 기존파일 skip 성공 / 미지원 backend LLM 에러 분기 검증
package agent

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

func TestScaffoldDDLMkdirError(t *testing.T) {
	// Creating "db" as a regular file makes os.MkdirAll(dbDir) fail.
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "db"), []byte("x"), 0o644); err != nil {
		t.Fatal(err)
	}
	ff := &features.FeaturesFile{Tables: map[string]features.TableDef{"users": {}}}
	var out bytes.Buffer
	if err := scaffoldDDL(dir, ff, nil, Config{}, &out); err == nil {
		t.Fatal("expected mkdir error when db is a file")
	}
}
