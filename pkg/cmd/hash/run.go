//ff:func feature=cli-hash type=command control=sequence
//ff:what Run — features.yaml 검증 후 SHA-256 해시를 specsDir/.yongol에 기록

package clihash

import (
	"crypto/sha256"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/cmd/featcheck"
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

// Run reads features.yaml inside specsDir, runs FT-* validation, computes
// its SHA-256 hash, and writes (or overwrites) the hash lock file at
// specsDir/.yongol. If validation produces ERROR-level diagnostics the
// hash is not written and an error is returned. WARNINGs are printed to
// out but do not block hash generation.
func Run(out io.Writer, specsDir string) error {
	featuresPath := filepath.Join(specsDir, "features.yaml")

	_, diags, err := featcheck.Run(featuresPath)
	if err != nil {
		return fmt.Errorf("features check: %w", err)
	}

	var errs []string
	for _, d := range diags {
		if d.Level == diagnostic.LevelError {
			errs = append(errs, d.Message)
		} else if d.Level == diagnostic.LevelWarning {
			fmt.Fprintf(out, "warning: %s\n", d.Message)
		}
	}
	if len(errs) > 0 {
		return fmt.Errorf("features validation failed:\n  %s", strings.Join(errs, "\n  "))
	}

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
