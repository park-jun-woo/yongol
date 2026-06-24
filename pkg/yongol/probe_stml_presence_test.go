//ff:func feature=orchestrator type=test control=iteration dimension=1
//ff:what TestProbeSTMLPresence — *.html glob 기반 Populated/Declared/Absent 3-상태 검증
package yongol

import (
	"os"
	"path/filepath"
	"testing"
)

func TestProbeSTMLPresence(t *testing.T) {
	root := t.TempDir()

	populated := filepath.Join(root, "with-html")
	if err := os.Mkdir(populated, 0o755); err != nil {
		t.Fatalf("mkdir populated: %v", err)
	}
	if err := os.WriteFile(filepath.Join(populated, "index.html"), []byte("<html></html>"), 0o644); err != nil {
		t.Fatalf("write html: %v", err)
	}

	empty := filepath.Join(root, "empty")
	if err := os.Mkdir(empty, 0o755); err != nil {
		t.Fatalf("mkdir empty: %v", err)
	}

	missing := filepath.Join(root, "missing")

	cases := []struct {
		name string
		dir  string
		want SSOTPresence
	}{
		{"with-html", populated, SSOTPopulated},
		{"empty-dir", empty, SSOTDeclared},
		{"missing-dir", missing, SSOTAbsent},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := probeSTMLPresence(tc.dir); got != tc.want {
				t.Fatalf("probeSTMLPresence(%q) = %v, want %v", tc.dir, got, tc.want)
			}
		})
	}
}
