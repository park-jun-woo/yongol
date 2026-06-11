//ff:func feature=gen-react type=test control=sequence
//ff:what writePackageJSON withAuthStore=true 시 zustand 의존성 포함 검증
package react

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func TestWritePackageJSON_WithAuthStore(t *testing.T) {
	dir := t.TempDir()
	if err := writePackageJSON(dir, true, false); err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(filepath.Join(dir, "package.json"))
	if err != nil {
		t.Fatal(err)
	}
	var pkg struct {
		Dependencies map[string]string `json:"dependencies"`
	}
	if err := json.Unmarshal(data, &pkg); err != nil {
		t.Fatalf("package.json is not valid JSON: %v", err)
	}
	if pkg.Dependencies["zustand"] == "" {
		t.Error("expected zustand dependency for the auth store")
	}
}
