//ff:func feature=agent type=test control=sequence
//ff:what TestScaffoldSQLc — 테이블 없음 nil / 기존파일 skip / 미지원 backend LLM 에러 / mkdir 에러 분기 검증
package agent

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

func TestScaffoldSQLcMkdirError(t *testing.T) {
	// Creating db/queries as a regular file makes os.MkdirAll fail.
	dir := t.TempDir()
	dbDir := filepath.Join(dir, "db")
	if err := os.MkdirAll(dbDir, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dbDir, "queries"), []byte("x"), 0o644); err != nil {
		t.Fatal(err)
	}
	ff := &features.FeaturesFile{Tables: map[string]features.TableDef{"users": {}}}
	var out bytes.Buffer
	if err := scaffoldSQLc(dir, ff, nil, Config{}, &out); err == nil {
		t.Fatal("expected mkdir error when db/queries is a file")
	}
}
