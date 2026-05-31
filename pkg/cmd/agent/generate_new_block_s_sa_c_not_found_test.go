//ff:func feature=agent type=test control=sequence
//ff:what TestGenerateNewBlock — SSaC 부재 skip / SSaC 존재+LLM 에러 skip 분기 검증
package agent

import (
	"bytes"
	"path/filepath"
	"strings"
	"testing"
)

func TestGenerateNewBlockSSaCNotFound(t *testing.T) {
	// No SSaC file under <specsDir>/service/*/ → early skip, returns false.
	dir := t.TempDir()
	abs := filepath.Join(dir, "api", "openapi.yaml")
	content := "paths: {}\n"
	var out bytes.Buffer
	if generateNewBlock(&out, Config{}, layerOpenAPI, "api/openapi.yaml", abs, &content, "ListUsers", "d", "p") {
		t.Fatal("expected false when SSaC file is missing")
	}
	if !strings.Contains(out.String(), "SSaC file not found") {
		t.Errorf("expected SSaC-not-found message, got: %q", out.String())
	}
}
