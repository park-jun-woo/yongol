//ff:func feature=features type=helper control=sequence
//ff:what writeHash — features.yaml SHA-256 해시를 specs/.yongol 에 기록

package features

import (
	"crypto/sha256"
	"fmt"
	"os"
	"path/filepath"
)

// writeHash computes SHA-256 of features data and writes specs/.yongol.
func writeHash(specsDir string, data []byte) error {
	hash := sha256.Sum256(data)
	content := fmt.Sprintf("hashes:\n  features.yaml: sha256:%x\n", hash)
	dest := filepath.Join(specsDir, ".yongol")
	if err := os.WriteFile(dest, []byte(content), 0o644); err != nil {
		return fmt.Errorf("write .yongol: %w", err)
	}
	return nil
}
