//ff:func feature=gen-react type=test control=sequence
//ff:what writeTailwindConfig fallback 기본 설정 생성 검증

package react

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestWriteTailwindConfig_Fallback(t *testing.T) {
	dir := t.TempDir()
	srcDir := filepath.Join(dir, "src")
	if err := os.MkdirAll(srcDir, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := writeTailwindConfig(dir, nil, nil); err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(filepath.Join(dir, "tailwind.config.js"))
	if err != nil {
		t.Fatal(err)
	}
	content := string(data)
	if !strings.Contains(content, "primary:") {
		t.Error("expected primary color in fallback config")
	}
	if !strings.Contains(content, "borderRadius:") {
		t.Error("expected borderRadius in fallback config")
	}
	if strings.Contains(content, "spacing:") {
		t.Error("fallback config should not have explicit spacing section")
	}
}
