//ff:func feature=cli-init type=util control=sequence
//ff:what copyFeaturesYAML — copies features.yaml into the specs directory

package cliinit

import (
	"fmt"
	"os"
	"path/filepath"
)

// copyFeaturesYAML copies the source features.yaml into
// <targetDir>/specs/features.yaml so it lives alongside the other SSOTs.
func copyFeaturesYAML(targetDir string, srcPath string) error {
	data, err := os.ReadFile(srcPath)
	if err != nil {
		return fmt.Errorf("read features.yaml: %w", err)
	}
	dest := filepath.Join(targetDir, "specs", "features.yaml")
	if err := os.WriteFile(dest, data, 0o644); err != nil {
		return fmt.Errorf("write features.yaml: %w", err)
	}
	return nil
}
