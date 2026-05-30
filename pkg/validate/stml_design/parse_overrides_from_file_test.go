//ff:func feature=validate type=test control=selection topic=stml-design
//ff:what TestParseOverridesFromFile — parseOverridesFromFile @override class 추출 분기 검증

package stml_design

import (
	"os"
	"path/filepath"
	"testing"
)

func TestParseOverridesFromFile(t *testing.T) {
	t.Run("missing file returns nil", func(t *testing.T) {
		if m := parseOverridesFromFile(filepath.Join(t.TempDir(), "nope.html")); m != nil {
			t.Errorf("expected nil for missing file, got %v", m)
		}
	})

	t.Run("override comment extracted", func(t *testing.T) {
		dir := t.TempDir()
		path := filepath.Join(dir, "page.html")
		os.WriteFile(path, []byte(
			`<main>
<!-- @override class="bg-neon-green p-8" -->
<div>Hello</div>
</main>`), 0o644)

		m := parseOverridesFromFile(path)
		if m == nil {
			t.Fatal("expected non-nil map")
		}
		if !m["bg-neon-green p-8"] {
			t.Errorf("expected override class extracted, got %v", m)
		}
	})
}
