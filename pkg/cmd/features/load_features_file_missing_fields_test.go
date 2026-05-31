//ff:func feature=features type=test control=iteration dimension=1
//ff:what TestLoadFeaturesFile — read/parse 에러·빈 features·필수필드 누락·성공 분기 검증
package features

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestLoadFeaturesFile_MissingFields(t *testing.T) {
	cases := []struct {
		name, body, want string
	}{
		{"op", "features:\n  - op: \"\"\n    path: POST /t\n    desc: d\n", "'op'"},
		{"path", "features:\n  - op: X\n    path: \"\"\n    desc: d\n", "'path'"},
		{"desc", "features:\n  - op: X\n    path: POST /t\n    desc: \"\"\n", "'desc'"},
	}
	for _, c := range cases {
		p := filepath.Join(t.TempDir(), "f.yaml")
		if err := os.WriteFile(p, []byte(c.body), 0o644); err != nil {
			t.Fatal(err)
		}
		_, err := loadFeaturesFile(p)
		if err == nil || !strings.Contains(err.Error(), c.want) {
			t.Errorf("%s: want %q error, got %v", c.name, c.want, err)
		}
	}
}
