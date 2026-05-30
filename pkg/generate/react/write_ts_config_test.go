//ff:func feature=gen-react type=test control=sequence
//ff:what writeTSConfig tsconfig.json 생성 내용·유효 JSON·에러경로 검증

package react

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func TestWriteTSConfig(t *testing.T) {
	dir := t.TempDir()
	if err := writeTSConfig(dir); err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(filepath.Join(dir, "tsconfig.json"))
	if err != nil {
		t.Fatal(err)
	}

	var cfg map[string]interface{}
	if err := json.Unmarshal(data, &cfg); err != nil {
		t.Fatalf("tsconfig.json is not valid JSON: %v", err)
	}
	opts, ok := cfg["compilerOptions"].(map[string]interface{})
	if !ok {
		t.Fatal("missing compilerOptions object")
	}
	if opts["strict"] != true {
		t.Error("expected strict=true")
	}
	if opts["jsx"] != "react-jsx" {
		t.Errorf("expected jsx=react-jsx, got %v", opts["jsx"])
	}
	paths, ok := opts["paths"].(map[string]interface{})
	if !ok {
		t.Fatal("missing paths object")
	}
	if _, ok := paths["@/*"]; !ok {
		t.Error("missing @/* path alias")
	}
}

func TestWriteTSConfigMissingDir(t *testing.T) {
	missing := filepath.Join(t.TempDir(), "absent", "dir")
	if err := writeTSConfig(missing); err == nil {
		t.Fatal("expected error writing into non-existent directory, got nil")
	}
}
