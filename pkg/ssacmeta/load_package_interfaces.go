//ff:func feature=ssacmeta type=loader control=iteration dimension=1
//ff:what LoadPackageInterfaces — ssac 루트의 pkg/*/interface.yaml 을 모두 로드

package ssacmeta

import (
	"fmt"
	"os"
	"path/filepath"
)

// LoadPackageInterfaces walks the given ssac repo root and loads every
// interface.yaml under pkg/*/interface.yaml. Returns a map keyed by the
// `package:` field.
func LoadPackageInterfaces(ssacRoot string) (map[string]*PackageInterface, error) {
	out := map[string]*PackageInterface{}
	pkgRoot := filepath.Join(ssacRoot, "pkg")
	entries, err := os.ReadDir(pkgRoot)
	if err != nil {
		if os.IsNotExist(err) {
			return out, nil
		}
		return nil, fmt.Errorf("ssacmeta: read pkg dir %s: %w", pkgRoot, err)
	}
	for _, e := range entries {
		if err := loadPackageInterfaceEntry(out, pkgRoot, e); err != nil {
			return nil, err
		}
	}
	return out, nil
}
