//ff:func feature=gen-react type=test control=sequence
//ff:what writeTailwindConfig postcss.config.js 생성 검증

package react

import (
	"os"
	"path/filepath"
	"testing"
)

func TestWriteTailwindConfig_PostcssEmitted(t *testing.T) {
	dir := t.TempDir()
	srcDir := filepath.Join(dir, "src")
	if err := os.MkdirAll(srcDir, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := writeTailwindConfig(dir, nil, nil); err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(filepath.Join(dir, "postcss.config.js")); os.IsNotExist(err) {
		t.Error("postcss.config.js not emitted")
	}
}
