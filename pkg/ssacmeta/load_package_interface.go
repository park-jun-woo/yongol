//ff:func feature=ssacmeta type=loader control=sequence
//ff:what LoadPackageInterface — 지정 경로의 interface.yaml 을 파싱

package ssacmeta

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// LoadPackageInterface reads and parses a single interface.yaml file.
// Returns nil (with nil error) if the path does not exist — callers treat
// absent files as "the package has no DB requirements".
func LoadPackageInterface(path string) (*PackageInterface, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, fmt.Errorf("ssacmeta: read %s: %w", path, err)
	}
	var iface PackageInterface
	if err := yaml.Unmarshal(data, &iface); err != nil {
		return nil, fmt.Errorf("ssacmeta: parse %s: %w", path, err)
	}
	iface.SourcePath = path
	return &iface, nil
}

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
		if !e.IsDir() {
			continue
		}
		ifacePath := filepath.Join(pkgRoot, e.Name(), "interface.yaml")
		iface, err := LoadPackageInterface(ifacePath)
		if err != nil {
			return nil, err
		}
		if iface == nil {
			continue
		}
		key := iface.Package
		if key == "" {
			key = e.Name()
		}
		out[key] = iface
	}
	return out, nil
}
