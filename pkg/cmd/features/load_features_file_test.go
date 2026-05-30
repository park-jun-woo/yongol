//ff:func feature=features type=test control=iteration dimension=1
//ff:what TestLoadFeaturesFile — read/parse 에러·빈 features·필수필드 누락·성공 분기 검증

package features

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestLoadFeaturesFile_ReadError(t *testing.T) {
	if _, err := loadFeaturesFile(filepath.Join(t.TempDir(), "nope.yaml")); err == nil ||
		!strings.Contains(err.Error(), "read features") {
		t.Fatalf("want read error, got %v", err)
	}
}

func TestLoadFeaturesFile_ParseError(t *testing.T) {
	p := filepath.Join(t.TempDir(), "f.yaml")
	if err := os.WriteFile(p, []byte("features: [oops"), 0o644); err != nil {
		t.Fatal(err)
	}
	if _, err := loadFeaturesFile(p); err == nil || !strings.Contains(err.Error(), "parse features") {
		t.Fatalf("want parse error, got %v", err)
	}
}

func TestLoadFeaturesFile_NoFeatures(t *testing.T) {
	p := filepath.Join(t.TempDir(), "f.yaml")
	if err := os.WriteFile(p, []byte("tables: {}\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	if _, err := loadFeaturesFile(p); err == nil || !strings.Contains(err.Error(), "no features") {
		t.Fatalf("want no-features error, got %v", err)
	}
}

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

func TestLoadFeaturesFile_Success(t *testing.T) {
	p := filepath.Join(t.TempDir(), "f.yaml")
	body := "features:\n  - op: X\n    path: POST /t\n    desc: d\n"
	if err := os.WriteFile(p, []byte(body), 0o644); err != nil {
		t.Fatal(err)
	}
	feats, err := loadFeaturesFile(p)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(feats) != 1 || feats[0].Op != "X" {
		t.Errorf("unexpected feats: %v", feats)
	}
}
