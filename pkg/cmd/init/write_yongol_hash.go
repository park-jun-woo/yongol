//ff:func feature=cli-init type=util control=sequence
//ff:what writeYongolHash — computes SHA-256 of features.yaml and writes specs/.yongol

package cliinit

import (
	"crypto/sha256"
	"fmt"
	"os"
	"path/filepath"
)

// writeYongolHash reads the features.yaml file at featuresPath, computes its
// SHA-256 hash, and writes the hash lock file at <targetDir>/specs/.yongol.
func writeYongolHash(targetDir string, featuresPath string) error {
	data, err := os.ReadFile(featuresPath)
	if err != nil {
		return fmt.Errorf("read features for hash: %w", err)
	}
	hash := sha256.Sum256(data)
	content := fmt.Sprintf("hashes:\n  features.yaml: sha256:%x\n", hash)

	dest := filepath.Join(targetDir, "specs", ".yongol")
	if err := os.WriteFile(dest, []byte(content), 0o644); err != nil {
		return fmt.Errorf("write .yongol: %w", err)
	}
	return nil
}
