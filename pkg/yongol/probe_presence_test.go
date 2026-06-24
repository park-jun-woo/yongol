//ff:func feature=orchestrator type=test control=iteration dimension=1
//ff:what TestProbePresence — 단일 파일 SSOT 존재/부재 → SSOTPopulated/SSOTAbsent 검증
package yongol

import (
	"os"
	"path/filepath"
	"testing"
)

func TestProbePresence(t *testing.T) {
	dir := t.TempDir()
	existing := filepath.Join(dir, "openapi.yaml")
	if err := os.WriteFile(existing, []byte("openapi: 3.0.0\n"), 0o644); err != nil {
		t.Fatalf("write fixture: %v", err)
	}

	cases := []struct {
		name string
		path string
		want SSOTPresence
	}{
		{"exists", existing, SSOTPopulated},
		{"not-exists", filepath.Join(dir, "missing.yaml"), SSOTAbsent},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := probePresence(tc.path); got != tc.want {
				t.Fatalf("probePresence(%q) = %v, want %v", tc.path, got, tc.want)
			}
		})
	}
}
