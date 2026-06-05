//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what writeComponent — 테스트용 <specsDir>/frontend/components/<name>.tsx 작성 헬퍼

package stml_openapi

import (
	"os"
	"path/filepath"
	"testing"
)

// writeComponent writes a <specsDir>/frontend/components/<name>.tsx file.
func writeComponent(t *testing.T, specsDir, name, body string) {
	t.Helper()
	dir := filepath.Join(specsDir, "frontend", "components")
	if err := os.MkdirAll(dir, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, name+".tsx"), []byte(body), 0o644); err != nil {
		t.Fatal(err)
	}
}
