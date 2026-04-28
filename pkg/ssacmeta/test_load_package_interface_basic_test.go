//ff:func feature=ssacmeta type=test control=sequence
//ff:what TestLoadPackageInterfaceBasic — 기본 interface.yaml 파싱 필드 검증

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
