//ff:func feature=ssacmeta type=test control=sequence
//ff:what TestLoadPackageInterfaceMissingReturnsNil — 파일 부재 시 (nil, nil) 반환

package ssacmeta

import (
	"path/filepath"
	"testing"
)

func TestLoadPackageInterfaceMissingReturnsNil(t *testing.T) {
	iface, err := LoadPackageInterface(filepath.Join(t.TempDir(), "missing.yaml"))
	if err != nil {
		t.Fatalf("missing file should not error: %v", err)
	}
	if iface != nil {
		t.Fatalf("iface should be nil for missing file")
	}
}
