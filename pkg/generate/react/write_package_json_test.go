//ff:func feature=gen-react type=test control=iteration dimension=1
//ff:what writePackageJSON package.json 생성 내용·유효 JSON·에러경로 검증
package react

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func TestWritePackageJSON(t *testing.T) {
	dir := t.TempDir()
	if err := writePackageJSON(dir, false, false); err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(filepath.Join(dir, "package.json"))
	if err != nil {
		t.Fatal(err)
	}

	var pkg struct {
		Private         bool              `json:"private"`
		Type            string            `json:"type"`
		Scripts         map[string]string `json:"scripts"`
		Dependencies    map[string]string `json:"dependencies"`
		DevDependencies map[string]string `json:"devDependencies"`
	}
	if err := json.Unmarshal(data, &pkg); err != nil {
		t.Fatalf("package.json is not valid JSON: %v", err)
	}

	if !pkg.Private {
		t.Error("expected private=true")
	}
	if pkg.Type != "module" {
		t.Errorf("expected type=module, got %q", pkg.Type)
	}
	for _, dep := range []string{"react", "react-dom", "react-router-dom", "@tanstack/react-query", "openapi-fetch"} {
		if _, ok := pkg.Dependencies[dep]; !ok {
			t.Errorf("missing dependency %q", dep)
		}
	}
	for _, dep := range []string{"vite", "typescript", "tailwindcss", "openapi-typescript"} {
		if _, ok := pkg.DevDependencies[dep]; !ok {
			t.Errorf("missing devDependency %q", dep)
		}
	}
	if pkg.Scripts["gen:api"] == "" {
		t.Error("missing gen:api script")
	}
	if _, ok := pkg.Dependencies["zustand"]; ok {
		t.Error("zustand must not be a dependency without the auth store")
	}
	if _, ok := pkg.Dependencies["lucide-react"]; ok {
		t.Error("lucide-react must not be a dependency without a sitemap data-icon")
	}
}
