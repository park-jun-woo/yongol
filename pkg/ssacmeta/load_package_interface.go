//ff:func feature=ssacmeta type=loader control=sequence
//ff:what LoadPackageInterface — 지정 경로의 interface.yaml 을 파싱

package ssacmeta

import (
	"fmt"
	"os"

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
