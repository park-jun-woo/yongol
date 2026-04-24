//ff:test feature=ssacmeta type=test
//ff:what LoadPackageInterface / LoadPackageInterfaces — 기본 파싱 경로 검증

package ssacmeta

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadPackageInterfaceBasic(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "interface.yaml")
	src := `version: 1
package: cache
description: sample
ports:
  - name: CacheSet
    description: set
    when: manifest.cache.backend == "postgres"
    used_by: [Set]
    query:
      cardinality: exec
      params:
        - { name: key, type: string }
        - { name: value, type: "[]byte" }
canonical_ddl: |
  CREATE TABLE fullend_cache(key TEXT PRIMARY KEY, value BYTEA NOT NULL);
`
	if err := os.WriteFile(path, []byte(src), 0o644); err != nil {
		t.Fatalf("write fixture: %v", err)
	}
	iface, err := LoadPackageInterface(path)
	if err != nil {
		t.Fatalf("LoadPackageInterface err: %v", err)
	}
	if iface == nil {
		t.Fatalf("iface nil")
	}
	if iface.Package != "cache" {
		t.Errorf("Package = %q, want cache", iface.Package)
	}
	if len(iface.Ports) != 1 {
		t.Fatalf("Ports len = %d, want 1", len(iface.Ports))
	}
	p := iface.Ports[0]
	if p.Name != "CacheSet" || p.Query.Cardinality != "exec" {
		t.Errorf("unexpected port: %+v", p)
	}
	if len(p.UsedBy) != 1 || p.UsedBy[0] != "Set" {
		t.Errorf("UsedBy = %v, want [Set]", p.UsedBy)
	}
	if iface.SourcePath != path {
		t.Errorf("SourcePath not populated")
	}
}

func TestLoadPackageInterfaceMissingReturnsNil(t *testing.T) {
	iface, err := LoadPackageInterface(filepath.Join(t.TempDir(), "missing.yaml"))
	if err != nil {
		t.Fatalf("missing file should not error: %v", err)
	}
	if iface != nil {
		t.Fatalf("iface should be nil for missing file")
	}
}

func TestLoadPackageInterfacesWalksPkgDir(t *testing.T) {
	root := t.TempDir()
	pkgDir := filepath.Join(root, "pkg")
	for _, name := range []string{"cache", "session", "notAnSsacPkg"} {
		if err := os.MkdirAll(filepath.Join(pkgDir, name), 0o755); err != nil {
			t.Fatalf("mkdir: %v", err)
		}
	}
	writeIface := func(name, pkg string) {
		src := "version: 1\npackage: " + pkg + "\nports: []\n"
		if err := os.WriteFile(filepath.Join(pkgDir, name, "interface.yaml"), []byte(src), 0o644); err != nil {
			t.Fatalf("write: %v", err)
		}
	}
	writeIface("cache", "cache")
	writeIface("session", "session")
	// notAnSsacPkg intentionally has no interface.yaml — should be skipped.

	ifaces, err := LoadPackageInterfaces(root)
	if err != nil {
		t.Fatalf("LoadPackageInterfaces err: %v", err)
	}
	if _, ok := ifaces["cache"]; !ok {
		t.Errorf("cache not loaded")
	}
	if _, ok := ifaces["session"]; !ok {
		t.Errorf("session not loaded")
	}
	if _, ok := ifaces["notAnSsacPkg"]; ok {
		t.Errorf("non-ssac pkg must be skipped")
	}
}

func TestEvaluateWhen(t *testing.T) {
	cases := []struct {
		expr string
		ctx  map[string]any
		want bool
	}{
		{"", nil, true},
		{"always", nil, true},
		{`manifest.cache.backend == "postgres"`, map[string]any{"cache": map[string]any{"backend": "postgres"}}, true},
		{`manifest.cache.backend == "postgres"`, map[string]any{"cache": map[string]any{"backend": "memory"}}, false},
		{`manifest.cache.backend == "postgres"`, map[string]any{}, false},
		{`manifest.backend.auth.refresh.enabled`, map[string]any{"backend": map[string]any{"auth": map[string]any{"refresh": map[string]any{"enabled": true}}}}, true},
		{`manifest.backend.auth.refresh.enabled`, map[string]any{"backend": map[string]any{"auth": map[string]any{"refresh": map[string]any{"enabled": false}}}}, false},
	}
	for _, c := range cases {
		got := EvaluateWhen(c.expr, c.ctx)
		if got != c.want {
			t.Errorf("EvaluateWhen(%q) = %v, want %v", c.expr, got, c.want)
		}
	}
}
