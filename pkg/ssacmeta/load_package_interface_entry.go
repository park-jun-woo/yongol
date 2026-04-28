//ff:func feature=ssacmeta type=loader control=sequence
//ff:what loadPackageInterfaceEntry — load a single pkg/<name>/interface.yaml into the out map

package ssacmeta

import (
	"os"
	"path/filepath"
)

// loadPackageInterfaceEntry loads a single pkg/<name>/interface.yaml entry
// into out, keyed by the `package:` field (falling back to the directory
// name). Non-directory entries and missing files are skipped silently.
func loadPackageInterfaceEntry(out map[string]*PackageInterface, pkgRoot string, e os.DirEntry) error {
	if !e.IsDir() {
		return nil
	}
	ifacePath := filepath.Join(pkgRoot, e.Name(), "interface.yaml")
	iface, err := LoadPackageInterface(ifacePath)
	if err != nil {
		return err
	}
	if iface == nil {
		return nil
	}
	key := iface.Package
	if key == "" {
		key = e.Name()
	}
	out[key] = iface
	return nil
}
