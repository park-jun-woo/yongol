//ff:func feature=cli-hash type=command control=sequence
//ff:what Run — reads features.yaml from specsDir, computes SHA-256, writes specsDir/.yongol

package clihash

import (
	"crypto/sha256"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// Run reads features.yaml inside specsDir, computes its SHA-256 hash, and
// writes (or overwrites) the hash lock file at specsDir/.yongol.
func Run(out io.Writer, specsDir string) error {
	featuresPath := filepath.Join(specsDir, "features.yaml")
	data, err := os.ReadFile(featuresPath)
	if err != nil {
		return fmt.Errorf("read features.yaml: %w", err)
	}

	hash := sha256.Sum256(data)
	content := fmt.Sprintf("hashes:\n  features.yaml: sha256:%x\n", hash)

	dest := filepath.Join(specsDir, ".yongol")
	if err := os.WriteFile(dest, []byte(content), 0o644); err != nil {
		return fmt.Errorf("write .yongol: %w", err)
	}

	fmt.Fprintf(out, "wrote %s (features.yaml sha256:%x)\n", dest, hash)
	return nil
}
