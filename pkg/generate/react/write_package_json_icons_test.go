//ff:func feature=gen-react type=test control=sequence
//ff:what writePackageJSON withIcons — lucide-react 의존성 조건부 추가 검증

package react

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func TestWritePackageJSON_Icons(t *testing.T) {
	dir := t.TempDir()
	if err := writePackageJSON(dir, false, true); err != nil {
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
	if _, ok := pkg.Dependencies["lucide-react"]; !ok {
		t.Error("expected lucide-react dependency with withIcons=true")
	}
}
