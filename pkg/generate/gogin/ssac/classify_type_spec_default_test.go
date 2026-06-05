//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestClassifyTypeSpecDefault — classifyResponseShapes 단위 테스트 (embedded struct vs schema alias 분류)

package ssac

import (
	"os"
	"path/filepath"
	"testing"
)

func TestClassifyTypeSpecDefault(t *testing.T) {
	dir := t.TempDir()
	src := `package api

type MapJSONResponse map[string]int

type EmbedNamedJSONResponse struct{ Named int }
`
	if err := os.WriteFile(filepath.Join(dir, "x.go"), []byte(src), 0o644); err != nil {
		t.Fatalf("write: %v", err)
	}
	shapes := classifyResponseShapes(dir)
	if _, ok := shapes["MapJSONResponse"]; ok {
		t.Error("map type should not be classified")
	}
	if _, ok := shapes["EmbedNamedJSONResponse"]; ok {
		t.Error("named field (non-anonymous) struct should be rejected")
	}
}
