//ff:func feature=agent type=test control=sequence
//ff:what TestScaffoldRego — 기존파일 skip / non-public 없음 skip / non-public 존재+미지원 backend LLM 에러 분기 검증
package agent

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

func TestScaffoldRegoMkdirError(t *testing.T) {
	// Creating "policy" as a regular file makes os.MkdirAll(policyDir) fail.
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "policy"), []byte("x"), 0o644); err != nil {
		t.Fatal(err)
	}
	ff := &features.FeaturesFile{Features: []features.Feature{{Op: "CreateUser", Public: false}}}
	var out bytes.Buffer
	if err := scaffoldRego(dir, ff, nil, Config{}, &out); err == nil {
		t.Fatal("expected mkdir error when policy is a file")
	}
}
